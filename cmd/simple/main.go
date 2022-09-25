package main

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/controllers"
	"github.com/Ktakuya332C/simple-clean-architecture/internal/repositories"
	"github.com/Ktakuya332C/simple-clean-architecture/internal/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	userRepository := repositories.NewUserRepositoryMemory()
	getUser := usecases.GetUser{UserRepository: userRepository}
	createUser := usecases.CreateUser{UserRepository: userRepository}
	getHandler := controllers.GetHandler{GetUser: getUser}
	createHandler := controllers.CreateHandler{CreateUser: createUser}

	engine := gin.Default()
	engine.GET("/get", getHandler.Handle)
	engine.POST("/create", createHandler.Handle)
	engine.Run(":8000")
}
