package sqlite

import (
	"errors"
	"fmt"
	"unsafe"
)

//go:linkname systemstack runtime.systemstack
func systemstack(fn func())

//go:linkname gostring runtime.gostring
func gostring(p *byte) string

// https://github.com/golang/go/blob/fc9f02c7aec81bcfcc95434d2529e0bb0bc03d66/src/runtime/env_posix.go#L85-L89.
func cstring(s string) unsafe.Pointer {
	p := make([]byte, len(s)+1)
	copy(p, s)
	return unsafe.Pointer(&p[0])
}

func errmsg(db unsafe.Pointer) error {
	var h unsafe.Pointer

	systemstack(func() {
		h = sqlite3_errmsg(db)
	})

	return errors.New(gostring((*byte)(h)))
}

type Stmt struct {
	handle unsafe.Pointer
	done   bool

	db *DB
}

func (stmt *Stmt) Done() bool {
	return stmt.done
}

func (stmt *Stmt) Step() error {
	var c int

	systemstack(func() {
		c = sqlite3_step(stmt.handle)
	})

	switch c {
	case 100: // row
		stmt.done = false

	case 101: // done
		stmt.done = true

	default:
		return errmsg(stmt.db.handle)
	}

	return nil
}

func (stmt *Stmt) Finalize() error {
	var c int

	systemstack(func() {
		c = sqlite3_finalize(stmt.handle)
	})

	if c != 0 {
		return errmsg(stmt.db.handle)
	}

	return nil
}

func (stmt *Stmt) BindInt(col int, val int) error {
	var c int

	systemstack(func() {
		c = sqlite3_bind_int(stmt.handle, col, val)
	})

	if c != 0 {
		return errmsg(stmt.db.handle)
	}

	return nil
}

func (stmt *Stmt) BindText(col int, val string) error {
	v := cstring(val)

	var c int

	systemstack(func() {
		c = sqlite3_bind_text(stmt.handle, col, v, -1, -1)
	})

	if c != 0 {
		return errmsg(stmt.db.handle)
	}

	return nil
}

func (stmt *Stmt) BindBlob(col int, val []byte) error {
	v := unsafe.Pointer(&val[0])

	var c int

	systemstack(func() {
		c = sqlite3_bind_text(stmt.handle, col, v, len(val), -1)
	})

	if c != 0 {
		return errmsg(stmt.db.handle)
	}

	return nil
}

func (stmt *Stmt) Reset() error {
	var c int

	systemstack(func() {
		c = sqlite3_reset(stmt.handle)
	})

	if c != 0 {
		return errmsg(stmt.db.handle)
	}

	return nil
}

func (stmt *Stmt) ColumnInt(col int) int {
	var v int

	systemstack(func() {
		v = sqlite3_column_int(stmt.handle, col)
	})

	return v
}

func (stmt *Stmt) ColumnText(col int) string {
	var v unsafe.Pointer

	systemstack(func() {
		v = sqlite3_column_text(stmt.handle, col)
	})

	return gostring((*byte)(v))
}

func (stmt *Stmt) ColumnBlob(col int) []byte {
	var s int
	var v unsafe.Pointer

	systemstack(func() {
		s = sqlite3_column_bytes(stmt.handle, col)
		v = sqlite3_column_blob(stmt.handle, col)
	})

	return unsafe.Slice((*byte)(v), s)

}

type DB struct {
	handle unsafe.Pointer
}

func (db *DB) Prepare(sql string) (*Stmt, error) {
	var stmt unsafe.Pointer

	s := cstring(sql)
	h := unsafe.Pointer(&stmt)

	var c int

	systemstack(func() {
		c = sqlite3_prepare_v2(db.handle, s, -1, h, 0)
	})

	if c != 0 {
		return nil, errmsg(db.handle)
	}

	return &Stmt{
		handle: stmt,

		db: db,
	}, nil
}

func (db *DB) Exec(sql string) error {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("prepare: %v", err)
	}

	err = stmt.Step()
	if err != nil {
		return fmt.Errorf("step: %v", err)
	}

	err = stmt.Finalize()
	if err != nil {
		return fmt.Errorf("finalize: %v", err)
	}

	return nil
}

func (db *DB) Close() {
	systemstack(func() {
		sqlite3_close(db.handle)
	})
}

func Open(name string) (*DB, error) {
	var db unsafe.Pointer

	n := cstring(name)
	h := unsafe.Pointer(&db)

	var c int

	systemstack(func() {
		c = sqlite3_open(n, h)
	})

	if c != 0 {
		return nil, errmsg(db)
	}

	return &DB{
		handle: db,
	}, nil
}
