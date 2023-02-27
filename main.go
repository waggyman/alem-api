package main
import (
    // "net/http"
    "github.com/gin-gonic/gin"
	"github.com/waggyman/alem-api/controllers"
	"github.com/waggyman/alem-api/models"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	// teachers side
	router.GET("/teachers", controllers.GetTeachers)
	router.POST("/teachers", controllers.StoreTeacher)
	router.GET("/teachers/:id", controllers.GetTeacherByID)
	router.DELETE("/teachers/:id", controllers.DeleteTeacherByID)

	router.Run("localhost:8000")
}