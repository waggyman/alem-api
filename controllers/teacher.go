package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waggyman/alem-api/models"
)

// type teacher struct {
// 	ID   string `json:"id"`
// 	Code string `json:"code"`
// 	Name string `json:"name"`
// }

// var teachers = []teacher{
// 	{ID: "1", Code: "002341", Name: "Junaedi M.A"},
// 	{ID: "2", Code: "002342", Name: "Fahrul S.A, M.A"},
// 	{ID: "3", Code: "002343", Name: "Akhmad Deedat M.A"},
// }

func GetTeachers(c *gin.Context) {
	teachers := make([]models.Teacher, 0)
	c.IndentedJSON(http.StatusOK, teachers)
}

func StoreTeacher(c *gin.Context) {
	var newTeacher models.Teacher
	if err := c.BindJSON(&newTeacher); err != nil {
		return
	}
	res := models.StoreTeacherMongo(newTeacher)
	fmt.Println(res)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "New Teacher Added"})
}

func GetTeacherByID(c *gin.Context) {
	// id := c.Param("id")
	var teacherFound models.Teacher
	// for _, teacher := range teachers {
	// 	if teacher.ID == id {
	// 		teacherFound = teacher
	// 	}
	// }

	// if (teacher{}) == teacherFound {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Teacher not found"})
	// 	return
	// }
	c.IndentedJSON(http.StatusOK, teacherFound)
}

func DeleteTeacherByID(c *gin.Context) {
	// id := c.Param("id")
	// var foundIndex int = -1
	// for i, teacher := range teachers {
	// 	if teacher.ID == id {
	// 		foundIndex = i
	// 	}
	// }

	// if foundIndex < 0 {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Teacher not found"})
	// 	return
	// }
	// newTeachers := make([]teacher, 0)
	// newTeachers = append(newTeachers, teachers[:foundIndex]...)
	// teachers = append(newTeachers, teachers[foundIndex+1:]...)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "AAAA"})
}
