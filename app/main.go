package main

import (
	"fmt"
	"github.com/tasuke/go-mux-task/controller"
	"github.com/tasuke/go-mux-task/db"
	"github.com/tasuke/go-mux-task/repositories"
	"github.com/tasuke/go-mux-task/route"
	"github.com/tasuke/go-mux-task/services"
	"github.com/tasuke/go-mux-task/utils/logic"
	"github.com/tasuke/go-mux-task/utils/validation"
	"log"
)

func main() {
	fmt.Println("Hello World")
	dbCon := db.NewDB()

	// logic
	authLogic := logic.NewAuthLogic()
	responseLogic := logic.NewResponseLogic()
	jwtLogic := logic.NewJWTLogic()

	// validation
	authValidation := validation.NewAuthValidation()
	// repository
	userRepository := repositories.NewUserRepository(dbCon)
	// service
	authService := services.NewAuthService(userRepository, authLogic, responseLogic, jwtLogic, authValidation)
	// controller
	authController := controller.NewAuthController(authService)
	// route
	authRouter := route.NewAuthRouter(authController)
	// main
	mainRouter := route.NewMainRouter(authRouter)

	if err := mainRouter.StartWebServer(); err != nil {
		log.Printf("error occured while starting server: %v", err)
	}
}
