package sqlserver

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewDBConnection() (*sql.DB, error) {
	server := "localhost"
	port := 1433
	user := "test5"
	password := "1111"
	database := "test5"

	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;",
		server, port, user, password, database)

	return sql.Open("sqlserver", connString)
}
