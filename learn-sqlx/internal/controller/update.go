package controller

import (
	"github.com/diyor200/learn-sqlx/internal/entity"
	"github.com/gin-gonic/gin"
)

func (c *Controller) UpdateUser(ctx *gin.Context) {
	body := entity.UserUpdateInput{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(400)
		return
	}
	c.crud.Update(body)
	ctx.JSON(200, gin.H{"message": "successfully updated!"})
	return
}
