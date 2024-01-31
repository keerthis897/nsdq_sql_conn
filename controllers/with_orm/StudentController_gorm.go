package controllers

import (
	"fmt"
	"log"
	"net/http"

	"nsdq_sql_conn/models"

	initializers "nsdq_sql_conn/initializers/with_orm"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Students
	err := initializers.DB.Find(&students).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Students
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON request"})
		return
	}

	if err := initializers.DB.Create(&student).Error; err != nil {
		// Log the error for debugging purposes
		log.Printf("Error creating student: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	c.JSON(http.StatusOK, student)
}
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Students

	if err := initializers.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	initializers.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := initializers.DB.Where("id = ?", id).Delete(&models.Students{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Student with ID %s deleted", id)})
}
