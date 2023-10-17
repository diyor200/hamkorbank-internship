package controller

import (
	"github.com/diyor200/gin-middleware-blogpost/internal/repository"
)

type Controller struct {
	r *repository.Repo
}

func NewController(r *repository.Repo) *Controller {
	return &Controller{r: r}
}
