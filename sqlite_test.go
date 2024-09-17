package sqlite

import (
	"testing"
)

func TestSqlite(t *testing.T) {
	db, err := Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("db: %+v\n", db)

	err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, data BLOB)")
	if err != nil {
		t.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO users (id, name, data) VALUES (?, ?, ?)")
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.BindInt(1, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.BindText(2, "Jim")
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.BindBlob(3, []byte{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.Step()
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.Reset()
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.BindInt(1, 2)
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.BindText(2, "Bob")
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.BindBlob(3, []byte{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.Step()
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.Finalize()
	if err != nil {
		t.Fatal(err)
	}

	stmt, err = db.Prepare("SELECT * FROM users")
	if err != nil {
		t.Fatal(err)
	}

	err = stmt.Step()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("done: %v\n", stmt.Done())

	t.Logf("%d %s %v\n", stmt.ColumnInt(0), stmt.ColumnText(1), stmt.ColumnBlob(2))

	db.Close()
}
