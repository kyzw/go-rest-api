package main

import (
	"go-reset-api/controller"
	"go-reset-api/db"
	"go-reset-api/repository"
	"go-reset-api/router"
	"go-reset-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
