package query

import _ "embed"

var (
	//go:embed create_schema.sql
	CreateSchema string
	//go:embed create_table.sql
	CreateTable string
	//go:embed user/user.insert_user.sql
	InsertUserRegistration string
)
