package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/julianaibiapina/projects/studies/golang-basic/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.GetUser)
	r.GET("/userByEmail/:userEmail", controller.GetUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
