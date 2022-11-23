package io_registerlogin

import (
	domain "STACK-GO/modules/io_loginregister/domain"
	intfService "STACK-GO/modules/io_loginregister/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	port intfService.Services
}

func NewHandler(port intfService.Services) *Handler {
	return &Handler{
		port: port,
	}
}

func (h *Handler) LoginCheck(c *gin.Context) {
	var input domain.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.port.Login(input.UName, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
