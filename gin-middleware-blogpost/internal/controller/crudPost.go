package controller

import (
	"errors"
	"fmt"
	"github.com/diyor200/gin-middleware-blogpost/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary		Create post
// @Description	Create new post
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param blog body entity.BlogInput true "enter your credentials"
// @Success		200		{object}	SuccessResponse
// @Failure		400,401	{object}	ErrorResponse
// @Security		ApiKeyAuth
// @Router			/action/create/post [POST]
func (c *Controller) CreatePost(ctx *gin.Context) {
	var err error
	userId, ok := ctx.Get("user_id")
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, ErrUnauthorized)
		return
	}
	body := entity.BlogInput{}

	if err = ctx.BindJSON(&body); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	postBody := entity.CreateBlogInput{PostBody: body.PostBody, PostTittle: body.PostTittle, UserID: userId.(int)}
	err = c.r.CreatePost(postBody)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Created post"})
	return
}

// @Summary		Get posts
// @Description	get all posts
// @Tags			posts
// @Produce		json
// @Success		200	{array}		entity.Blog
// @Failure		400	{object}	ErrorResponse
// @Router			/posts [GET]
func (c *Controller) GetPosts(ctx *gin.Context) {
	posts, err := c.r.GetPosts()
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, gin.H{"posts": posts})
	return
}

// @Summary		Get post
// @Description	get post by id
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			post_id	path		int	true	"Post id"
// @Success		200		{object}	entity.Blog
// @Failure		400		{object}	ErrorResponse
// @Router			/posts/{post_id} [GET]
func (c *Controller) GetPost(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("post_id"))
	fmt.Println("GetPost - post_id = ", postId)
	if err != nil {
		errorResponse(ctx, 400, err)
		return
	}
	post, err := c.r.GetPost(postId)
	if err != nil || len(post) == 0 {
		errorResponse(ctx, http.StatusBadRequest, errors.New("not found"))
		return
	}

	ctx.JSON(200, gin.H{"blog": post})
	return
}

// @Summary		Edit post
// @Description	Edit post
// @Tags			posts
// @Accept			json
// @Produce		json
// @Param			post	body		entity.BlogInput	true	"Post to update"
// @Success		200		{object}	SuccessResponse
// @Failure		400		{object}	ErrorResponse
// @Router			/action/edit/post [POST]
// @Security		ApiKeyAuth
func (c *Controller) EditPost(ctx *gin.Context) {
	var input entity.BlogInput
	var err error
	userId, ok := ctx.Get("user_id")
	fmt.Println("EditPost userId = ", userId)
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, ErrUnauthorized)
		return
	}
	if err = ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = c.r.EditPost(input, userId.(int))
	log.Println("err = c.r.EditPost(input, userId.(int)) = ", err)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Updated post"})
	return
}

// @Summary		Delete post
// @Description	Delete one post
// @Tags			posts
// @Param			post_id	path		int	true	"Post id"
// @Success		200		{object}	SuccessResponse
// @Failure		400		{object}	ErrorResponse
// @Router			/action/delete/post/{post_id} [POST]
// @Security		ApiKeyAuth
func (c *Controller) DeletePost(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("post_id"))
	userId, ok := ctx.Get("user_id")
	if !ok {
		errorResponse(ctx, http.StatusUnauthorized, ErrUnauthorized)
		return
	}

	err = c.r.DeletePost(postId, userId.(int))
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "post deleted"})
}

func errorResponse(ctx *gin.Context, status int, err error) {
	ctx.AbortWithStatusJSON(status, map[string]string{"error": fmt.Sprintf("%v", err)})
}
