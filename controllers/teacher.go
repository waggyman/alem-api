package controllers
import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type teacher struct {
	ID 		string `json:"id"`
	Code	string `json:"code"`
	Name 	string `json:"name"`
}

var teachers = []teacher{
	{ID: "1", Code: "002341", Name: "Junaedi M.A"},
	{ID: "2", Code: "002342", Name: "Fahrul S.A, M.A"},
	{ID: "3", Code: "002343", Name: "Akhmad Deedat M.A"},
}

func GetTeachers (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teachers)
}

func StoreTeacher (c *gin.Context) {
	var newTeacher teacher

	if err := c.BindJSON(&newTeacher); err != nil {
		return
	}

	teachers = append(teachers, newTeacher)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "New Teacher Added", "data": newTeacher})
}

func GetTeacherByID (c *gin.Context) {
	id := c.Param("id")
	var teacherFound teacher
	for _, teacher := range	teachers {
		if teacher.ID == id {
			teacherFound = teacher
		}
	}
	
	if (teacher{}) == teacherFound {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Teacher not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, teacherFound)
}

func DeleteTeacherByID (c *gin.Context) {
	id := c.Param("id")
	var foundIndex int = -1
	for i, teacher := range teachers {
		if teacher.ID == id {
			foundIndex = i
		}
	}

	if (foundIndex < 0) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Teacher not found"})
		return		
	}
	newTeachers := make([]teacher, 0)
	newTeachers = append(newTeachers, teachers[:foundIndex]...)
	teachers = append(newTeachers, teachers[foundIndex + 1:]...)
	c.IndentedJSON(http.StatusOK, teachers)
}