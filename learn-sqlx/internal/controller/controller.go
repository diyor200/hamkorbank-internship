package controller

import "github.com/diyor200/learn-sqlx/crud"

type Controller struct {
	crud *crud.Crud
}

func NewController(crud *crud.Crud) *Controller {
	return &Controller{
		crud: crud,
	}
}
