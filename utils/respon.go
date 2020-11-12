package utils

import (
	"mongodb-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func HandleSuccess(c *gin.Context, data interface{}) {
	var res = model.ResponWrapper{
		Success: true,
		Message: "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func HandleError(c *gin.Context, status int, message string) {
	var res = model.ResponWrapper{
		Success: false,
		Message: message,
	}
	c.JSON(status, res)
}