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
	// User
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	// Task
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	// pass struct to constructor for DI
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
