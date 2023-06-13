package main

import (
	"go-reset-api/controller"
	"go-reset-api/db"
	"go-reset-api/repository"
	"go-reset-api/router"
	"go-reset-api/usecase"
	"go-reset-api/validator"
)

func main() {
	db := db.NewDB()
	// User
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	// Task
	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	// pass struct to constructor for DI
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
