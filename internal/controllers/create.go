package controllers

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/entities"
	"github.com/Ktakuya332C/simple-clean-architecture/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateHandler struct {
	CreateUser usecases.CreateUser
}

func (inst *CreateHandler) Handle(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := inst.CreateUser.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update success"})
}
