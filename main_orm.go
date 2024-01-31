package main

import (
	"fmt"
	initializers "nsdq_sql_conn/initializers/with_orm"

	"nsdq_sql_conn/models"
	routes "nsdq_sql_conn/routes/with_orm"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.DB.AutoMigrate(&models.Students{})

}

func main() {

	router := gin.Default()

	routes.SetupRoutes(router)

	fmt.Println("Server listening on port 8080...")
	router.Run(":8080")

}
