module github.com/cdvelop/file

go 1.20

require (
	github.com/cdvelop/api v0.0.15
	github.com/cdvelop/cutkey v0.6.0
	github.com/cdvelop/dbtools v0.0.41 // indirect
	github.com/cdvelop/input v0.0.34
	github.com/cdvelop/testools v0.0.12
)

require (
	github.com/cdvelop/objectdb v0.0.66 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)

require (
	github.com/cdvelop/gotools v0.0.42
	github.com/cdvelop/model v0.0.53
	github.com/cdvelop/output v0.0.3
	github.com/cdvelop/sqlite v0.0.41
	github.com/cdvelop/timeserver v0.0.1
	github.com/cdvelop/unixid v0.0.3
	github.com/gabriel-vasile/mimetype v1.4.2
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/unixid => ../unixid

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/api => ../api

replace github.com/cdvelop/sqlite => ../sqlite

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/objectdb => ../objectdb

replace github.com/cdvelop/dbtools => ../dbtools

replace github.com/cdvelop/output => ../output

replace github.com/cdvelop/testools => ../testools
