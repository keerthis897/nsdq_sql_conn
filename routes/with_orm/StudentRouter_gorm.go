package routes

import (
	controllers "nsdq_sql_conn/controllers/with_orm"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/students", controllers.GetStudents)
	router.POST("/students/create", controllers.CreateStudent)
	router.POST("/students/:id/update", controllers.UpdateStudent)
	router.DELETE("/students/:id", controllers.DeleteStudent)
}
