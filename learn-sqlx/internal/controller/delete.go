package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type deleteInput struct {
	ID int
}

func (c *Controller) Delete(ctx *gin.Context) {
	input := deleteInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println("=== user id for delete === ", input.ID)

	c.crud.DeleteUser(input.ID)

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
	return
}
