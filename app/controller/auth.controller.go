package controller

import (
	"github.com/tasuke/go-mux-task/services"
	"net/http"
)

type IAuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	as services.IAuthService
}

func NewAuthController(as services.IAuthService) IAuthController {
	return &authController{as}
}

func (ac authController) SignIn(w http.ResponseWriter, r *http.Request) {
	user, err := ac.as.SignIn(w, r)
	if err != nil {
		return
	}

	ac.as.SendAuthResponse(w, &user, http.StatusOK)
}

func (ac authController) SignUp(w http.ResponseWriter, r *http.Request) {
	user, err := ac.as.SignUp(w, r)
	if err != nil {
		return
	}

	ac.as.SendAuthResponse(w, &user, http.StatusCreated)
}
