package sqlite

import "unsafe"

//go:cgo_import_dynamic sqlite3_open sqlite3_open "libsqlite3.dylib"
func sqlite3_open(name unsafe.Pointer, handle unsafe.Pointer) int

//go:cgo_import_dynamic sqlite3_close sqlite3_close "libsqlite3.dylib"
func sqlite3_close(handle unsafe.Pointer)

//go:cgo_import_dynamic sqlite3_prepare_v2 sqlite3_prepare_v2 "libsqlite3.dylib"
func sqlite3_prepare_v2(handle unsafe.Pointer, sql unsafe.Pointer, bytes int, statement unsafe.Pointer, tail uintptr) int

//go:cgo_import_dynamic sqlite3_step sqlite3_step "libsqlite3.dylib"
func sqlite3_step(handle unsafe.Pointer) int

//go:cgo_import_dynamic sqlite3_errmsg sqlite3_errmsg "libsqlite3.dylib"
func sqlite3_errmsg(handle unsafe.Pointer) unsafe.Pointer

//go:cgo_import_dynamic sqlite3_finalize sqlite3_finalize "libsqlite3.dylib"
func sqlite3_finalize(handle unsafe.Pointer) int

//go:cgo_import_dynamic sqlite3_bind_int sqlite3_bind_int "libsqlite3.dylib"
func sqlite3_bind_int(handle unsafe.Pointer, column int, value int) int

//go:cgo_import_dynamic sqlite3_bind_text sqlite3_bind_text "libsqlite3.dylib"
func sqlite3_bind_text(handle unsafe.Pointer, column int, value unsafe.Pointer, length int, callback int) int

//go:cgo_import_dynamic sqlite3_bind_blob sqlite3_bind_blob "libsqlite3.dylib"
func sqlite3_bind_blob(handle unsafe.Pointer, column int, value unsafe.Pointer, length int, callback int) int

//go:cgo_import_dynamic sqlite3_reset sqlite3_reset "libsqlite3.dylib"
func sqlite3_reset(handle unsafe.Pointer) int

//go:cgo_import_dynamic sqlite3_column_int sqlite3_column_int "libsqlite3.dylib"
func sqlite3_column_int(handle unsafe.Pointer, column int) int

//go:cgo_import_dynamic sqlite3_column_text sqlite3_column_text "libsqlite3.dylib"
func sqlite3_column_text(handle unsafe.Pointer, column int) unsafe.Pointer

//go:cgo_import_dynamic sqlite3_column_blob sqlite3_column_blob "libsqlite3.dylib"
func sqlite3_column_blob(handle unsafe.Pointer, column int) unsafe.Pointer

//go:cgo_import_dynamic sqlite3_column_bytes sqlite3_column_bytes "libsqlite3.dylib"
func sqlite3_column_bytes(handle unsafe.Pointer, column int) int
