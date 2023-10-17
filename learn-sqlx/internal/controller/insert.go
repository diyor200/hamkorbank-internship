package controller

import (
	"github.com/diyor200/learn-sqlx/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (c *Controller) InsertUser(ctx *gin.Context) {
	input := entity.UserInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println(input)
	c.crud.InsertUser(input)
	ctx.JSONP(http.StatusCreated, gin.H{"message": "successfully created"})
	return
}
