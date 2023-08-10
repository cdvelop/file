module github.com/cdvelop/file

go 1.20

require (
	github.com/cdvelop/api v0.0.0-00010101000000-000000000000
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/dbtools v0.0.26
	github.com/cdvelop/input v0.0.16
	github.com/cdvelop/testools v0.0.0-00010101000000-000000000000
)

require (
	github.com/cdvelop/objectdb v0.0.38 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	golang.org/x/net v0.13.0 // indirect
)

require (
	github.com/cdvelop/gotools v0.0.25
	github.com/cdvelop/model v0.0.36
	github.com/cdvelop/sqlite v0.0.34
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

replace github.com/cdvelop/testools => ../testools
