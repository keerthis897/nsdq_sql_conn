package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	server := "localhost"
	port := 1433
	user := "kk"
	password := "1111"
	database := "test2"

	connString := fmt.Sprintf("server=%s;port=%d;userid=%s;password=%s;database=%s;",
		server, port, user, password, database)

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer conn.Close()

	ctx := context.Background()

	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	fmt.Println("Connected to SQL Server!")

}
