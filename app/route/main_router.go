package route

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type IMainRouter interface {
	setupRouter() *mux.Router
	StartWebServer() error
}
type mainRouter struct {
	authR IAuthRouter
}

func NewMainRouter(authR IAuthRouter) IMainRouter {
	return &mainRouter{authR}
}

func (mr *mainRouter) setupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	mr.authR.SetupAuthRouter(router)
	return router
}

func (mr *mainRouter) StartWebServer() error {
	fmt.Println("Starting server...")

	return http.ListenAndServe(":8080", mr.setupRouter())
}
