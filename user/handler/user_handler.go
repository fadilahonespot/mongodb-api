package handler

import (
	"fmt"
	"mongodb-api/model"
	"mongodb-api/user"
	"net/http"
	"mongodb-api/utils"

	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	userUsecase user.UserUsecase
}

func CreateUserHandler(r *gin.Engine, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}

	r.GET("/user", userHandler.findAllUser)
	r.POST("/user", userHandler.addUser)
	r.GET("/user/:id", userHandler.findUser)
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.DeleteUser)
}

func (e *UserHandler) addUser(c *gin.Context) {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("Error bind data %v \n", err)
		return
	}
	result, err := e.userUsecase.InsertUser(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, result)
}

func (e *UserHandler) findAllUser(c *gin.Context) {
	user, err := e.userUsecase.FindAllUser()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, user)
}

func (e *UserHandler) findUser(c *gin.Context) {
	id := c.Param("id")
	user, err := e.userUsecase.FindUser(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, "User Not Found")
		return
	}
	utils.HandleSuccess(c, user)
}

func (e *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		fmt.Printf("error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, "Oppss internal server error")
		return
	}
	result, err := e.userUsecase.UpdateUser(id, &user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccess(c, result)
}

func (e *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := e.userUsecase.DeleteUser(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.HandleSuccess(c, "Success Deleted Data")
}