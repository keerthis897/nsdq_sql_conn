package main

import (
	"context"
	"fmt"
	"log"
	"nsdq_sql_conn/sqlserver"
)

func main() {
	conn, err := sqlserver.NewDBConnection()
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer conn.Close()

	// Perform SQL operations
	ctx := context.Background()
	err = sqlserver.CreateTableIfNotExists(ctx, conn)
	if err != nil {
		log.Fatal("Error creating table: ", err.Error())
	}

	rowsAffected, err := sqlserver.InsertData(ctx, conn, "John4")
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}
	fmt.Printf("%d row(s) inserted.\n", rowsAffected)
}
