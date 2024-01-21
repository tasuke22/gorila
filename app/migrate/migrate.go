package main

import (
	"fmt"
	"github.com/tasuke/go-mux-task/db"
	"github.com/tasuke/go-mux-task/models"
)

func main() {
	dbCon := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbCon)
	dbCon.AutoMigrate(&models.User{}, &models.Todo{})
}
