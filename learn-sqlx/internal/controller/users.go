package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) GetUsers(ctx *gin.Context) {
	users, err := c.crud.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
	return
}
