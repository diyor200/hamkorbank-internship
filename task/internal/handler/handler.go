package handler

import (
	"absapi/internal/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetUserMap(ctx *gin.Context) {
	userID := ctx.Param("persId")
	persID, err := strconv.Atoi(userID)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user, err := h.repo.GetUser(persID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			userFromABS, err := h.GetUserFromABSMAp(persID)
			if err != nil {
				ctx.JSON(400, gin.H{"message": err.Error()})
				return
			}

			u, err := h.repo.InsertUserFromMap(userFromABS)
			if err != nil {
				ctx.JSON(400, gin.H{"message": err.Error()})
			}
			ctx.JSON(200, &u)
			return
		}
		ctx.JSON(500, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, &user)
	return
}

//func (h *Handler) GetUser(ctx *gin.Context) {
//	userID := ctx.Param("persId")
//	fmt.Println(userID)
//	persID, err := strconv.Atoi(userID)
//	if err != nil {
//		ctx.JSON(400, gin.H{"message": err.Error()})
//		return
//	}
//	user, err := h.repo.GetUser(persID)
//	if err != nil {
//		if errors.Is(err, repository.ErrNotFound) {
//			userFromABS, err := h.GetUserFromABS(persID)
//			if err != nil {
//				ctx.JSON(400, gin.H{"message": err.Error()})
//				return
//			}
//			if len(userFromABS.ResponseBody) == 0 {
//				ctx.JSON(404, gin.H{"message": "User not found"})
//				return
//			}
//			h.repo.InsertUser(userFromABS)
//			res, err := h.repo.GetUser(persID)
//			if err != nil {
//
//				ctx.JSON(400, gin.H{"message": err.Error()})
//				return
//			}
//			ctx.JSON(200, &res)
//			return
//		}
//		ctx.JSON(500, gin.H{"message": err.Error()})
//		return
//	}
//	ctx.JSON(200, &user)
//	return
//}
