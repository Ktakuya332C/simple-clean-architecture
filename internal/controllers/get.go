package controllers

import (
	"github.com/Ktakuya332C/simple-clean-architecture/internal/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetHandler struct {
	GetUser usecases.GetUser
}

func (inst *GetHandler) Handle(c *gin.Context) {
	userID := c.Query("user_id")
	user, err := inst.GetUser.Get(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"FirstName": user.FirstName, "LastName": user.LastName})
}
