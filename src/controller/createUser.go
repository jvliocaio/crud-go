package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jvliocaio/crud-go/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("moscou")
	c.JSON(int(err.Code), err)
}
