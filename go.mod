module github.com/cdvelop/file

go 1.20

require (
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/input v0.0.15
)

require (
	github.com/cdvelop/dbtools v0.0.25 // indirect
	github.com/cdvelop/objectdb v0.0.37 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.12.0 // indirect
)

require (
	github.com/cdvelop/api v0.0.6
	github.com/cdvelop/gotools v0.0.18
	github.com/cdvelop/model v0.0.34
	github.com/cdvelop/sqlite v0.0.29
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/gabriel-vasile/mimetype v1.4.2
	golang.org/x/text v0.11.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/api => ../api

replace github.com/cdvelop/sqlite => ../sqlite

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/dbtools => ../dbtools
