package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	server := "localhost"
	port := 1433
	user := "test5"
	password := "1111"
	database := "test5"

	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;",
		server, port, user, password, database)

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	fmt.Println("Connected to SQL Server!")

	_, err = conn.Exec(`
	IF NOT EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[dbo].[students]') AND type in (N'U'))
	BEGIN
		CREATE TABLE students (
			ID INT PRIMARY KEY IDENTITY,
			Name NVARCHAR(50)
		)
	END`)
	if err != nil {
		log.Fatal("Error creating table: ", err.Error())
	}

	result, err := conn.Exec("INSERT INTO students (name) VALUES ('John6')")
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error getting rows affected: ", err.Error())
	}
	fmt.Printf("%d row(s) inserted.\n", rowsAffected)

	result, err = conn.Exec("DELETE FROM students WHERE Name = 'Johna'")
	if err != nil {
		log.Fatal("Error deleting data: ", err.Error())
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Fatal("Error getting rows affected: ", err.Error())
	}
	fmt.Printf("%d row(s) deleted.\n", rowsAffected)

	result, err = conn.Exec("UPDATE students SET Name = 'Johnny' WHERE Name = 'John'")
	if err != nil {
		log.Fatal("Error updating data: ", err.Error())
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Fatal("Error getting rows affected: ", err.Error())
	}
	fmt.Printf("%d row(s) updated.\n", rowsAffected)
}
