package route

import (
	"github.com/gorilla/mux"
	"github.com/tasuke/go-mux-task/controller"
)

type IAuthRouter interface {
	SetupAuthRouter(router *mux.Router)
}

type authRouter struct {
	ac controller.IAuthController
}

func NewAuthRouter(ac controller.IAuthController) IAuthRouter {
	return &authRouter{ac}
}

func (ar authRouter) SetupAuthRouter(router *mux.Router) {
	router.HandleFunc("/signin", ar.ac.SignIn).Methods("POST")
	router.HandleFunc("/signup", ar.ac.SignUp).Methods("POST")
}
