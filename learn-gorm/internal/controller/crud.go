package controller

import (
	"github.com/diyor200/learn-gorm/internal/entity"
	"github.com/diyor200/learn-gorm/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (c *Controller) Create(ctx *gin.Context) {
	input := entity.UserInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := service.Create(c.db, input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "successfully created"})
	return
}

func (c *Controller) Read(ctx *gin.Context) {
	var users []entity.User
	users, err := service.Read(c.db)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, &users)
	return
}
func (c *Controller) Update(ctx *gin.Context) {
	input := entity.UserInput{}
	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := service.Update(c.db, input)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, &user)
	return
}

func (c *Controller) Delete(ctx *gin.Context) {
	value := ctx.Param("id")
	id, err := strconv.Atoi(value)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println("controller user_id=====", id)
	err = service.Delete(c.db, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, gin.H{"message": "deleted"})
	return
}
