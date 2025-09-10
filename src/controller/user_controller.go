package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	rest_error "github.com/julianaibiapina/projects/studies/golang-basic/src/configuration/rest_errors"
	"github.com/julianaibiapina/projects/studies/golang-basic/src/controller/model/request"
)

func GetUser(context *gin.Context) {
}

func GetUserByEmail(context *gin.Context) {
}

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		msg := fmt.Sprintf("There are errors in the passed fields, error=%s", err.Error())
		err := rest_error.NewBadRequestError(msg)
		context.JSON(err.Code, err)
		return
	}

	fmt.Println(userRequest)
}

func UpdateUser(context *gin.Context) {
}

func DeleteUser(context *gin.Context) {
}
