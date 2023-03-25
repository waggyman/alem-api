package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waggyman/alem-api/models"
)

type SetSubjectParams struct {
	Assign   []string `bson:"set,omitempty"`
	Unassign []string `bson:"unset,omitempty"`
	Set      []string `bson:"set,omitempty"`
}

func GetTeachers(c *gin.Context) {
	// teachers := make([]models.Teacher, 0)
	teachers := models.ListTeacherMongo()
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

func FindTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacherFound models.Teacher
	res := models.FindTeacherByIdMongo(id)
	teacherFound = res
	fmt.Println(res)
	c.IndentedJSON(http.StatusOK, teacherFound)
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var payload models.Teacher
	if err := c.BindJSON(&payload); err != nil {
		return
	}
	updatedTeacher := models.UpdateTeacherById(id, payload)
	c.IndentedJSON(http.StatusOK, updatedTeacher)
}

func RemoveTeacher(c *gin.Context) {
	id := c.Param("id")
	res := models.DeleteTeacherByID(id)
	message := "Successfully Deleted"
	if !res {
		message = "Unsuccessfully Deleted"
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": message})
}

func SetSubjectToTeacher(c *gin.Context) {
	id := c.Param("id")
	var payload SetSubjectParams
	if err := c.BindJSON(&payload); err != nil {
		return
	}
	currentTeacher := models.FindTeacherByIdMongo(id)
	// if (len(payload.Assign) > 0) {
	if len(payload.Set) > 0 {
		for _, v := range payload.Set {
			foundSubject := models.GetSubjectByID(v)
			if (foundSubject == models.Subject{}) {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"ERROR": "Subject Not Found"})
				return
			}
		}
	} else {
		for _, v := range payload.Assign {
			foundSubject := models.GetSubjectByID(v)
			fmt.Println(v)
			if (foundSubject == models.Subject{}) {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"ERROR": "Subject Not Found"})
				return
			}
		}

		for _, v := range payload.Unassign {
			foundSubject := models.GetSubjectByID(v)
			if (foundSubject == models.Subject{}) {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"ERROR": "Subject Not Found"})
				return
			}
		}
	}
	// }
	fmt.Println(currentTeacher)
	c.IndentedJSON(http.StatusOK, payload)
}
