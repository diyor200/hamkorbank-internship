package controller

import "errors"

var ErrUnauthorized = errors.New("Sign In first")

type ErrorResponse struct {
	Message string `json:"error"`
}
type SuccessResponse struct {
	Message string `json:"message"`
}
