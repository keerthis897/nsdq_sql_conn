package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	initializers "nsdq_sql_conn/initializers/without_orm"
)

func CreateStudentTable(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	conn := initializers.GetDB()
	if conn == nil {
		http.Error(w, "Db connection is nil", http.StatusInternalServerError)
		return
	}

	err := initializers.CreateTableIfNotExists(ctx, conn)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating table: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Table students created successfully if it didn't exist already.")
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Retrieve the database connection
	db := initializers.GetDB()
	if db == nil {
		http.Error(w, "Db connection is nil", http.StatusInternalServerError)
		return
	}

	// Query to retrieve students
	rows, err := db.Query("SELECT ID, Name FROM students")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Slice to hold retrieved students
	var students []Student

	// Iterate over rows
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error iterating over rows: %v", err), http.StatusInternalServerError)
		return
	}

	// Encode students slice to JSON
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	conn := initializers.GetDB()
	if conn == nil {
		log.Fatal("DB connection is nil", http.StatusInternalServerError)
	}
	defer conn.Close()

	result, err := conn.Exec("INSERT INTO students (Name) VALUES ('John6')")
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error in getting rows affected: ", err.Error())
	}
	fmt.Fprintf(w, "%d row(s) inserted.\n", rowsAffected)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	conn := initializers.GetDB()
	if conn == nil {
		log.Fatal("Database connection is nil", http.StatusInternalServerError)
	}
	defer conn.Close()

	result, err := conn.Exec("UPDATE students SET Name = 'Johnny' WHERE Name = 'John'")
	if err != nil {
		log.Fatal("Error updating data: ", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error getting rows affected: ", err.Error())
	}
	fmt.Fprintf(w, "%d row(s) updated.\n", rowsAffected)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	conn := initializers.GetDB()
	if conn == nil {
		log.Fatal("Database connection is nil", http.StatusInternalServerError)
	}
	defer conn.Close()

	result, err := conn.Exec("DELETE FROM students WHERE Name = 'Johna'")
	if err != nil {
		log.Fatal("Error deleting data: ", err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error getting rows affected: ", err.Error())
	}
	fmt.Fprintf(w, "%d row(s) deleted.\n", rowsAffected)
}
