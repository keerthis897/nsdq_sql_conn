package main

import (
	"fmt"
	"net/http"

	initializers "nsdq_sql_conn/initializers/without_orm"
	routes "nsdq_sql_conn/routes/without_orm"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

}

func main() {

	routes.SetupRoutes()

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
