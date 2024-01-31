package routes

import (
	"net/http"

	controllers "nsdq_sql_conn/controllers/without_orm"
)

func SetupRoutes() {
	http.HandleFunc("/students", controllers.GetStudents)
	http.HandleFunc("/students/create", controllers.CreateStudent)
	http.HandleFunc("/students/update", controllers.UpdateStudent)
	http.HandleFunc("/students/delete", controllers.DeleteStudent)
	http.HandleFunc("/students/create_table", controllers.CreateStudentTable) // New route for creating table
}
