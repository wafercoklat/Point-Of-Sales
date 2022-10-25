package user

import (
	domain "STACK-GO/modules/user/domain"
	intfService "STACK-GO/modules/user/interface"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	port intfService.Services
}

func NewHandler(route *gin.Engine, port intfService.Services) {
	http := Handler{
		port: port,
	}

	// Routes
	route.GET("/user/:id/", http.FindByID)
	route.GET("/user/list/", http.List)
	route.POST("/user/add/", http.Create)
	route.PUT("/user/update/:id/", http.Update)
	route.DELETE("/user/delete/:id/", http.Delete)
}

func (h *Handler) FindByID(c *gin.Context) {
	val := h.port.FindByID(c.Param("id"))
	c.JSON(200, val)
}

func (h *Handler) List(c *gin.Context) {
	val := h.port.List()
	c.JSON(200, val)
}

func (h *Handler) Create(c *gin.Context) {
	var request domain.User
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val := h.port.Create(request)
	c.JSON(200, val)
}

func (h *Handler) Update(c *gin.Context) {
	var request domain.User
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val := h.port.Update(c.Param("id"), request)
	c.JSON(200, val)
}

func (h *Handler) Delete(c *gin.Context) {
	val := h.port.Delete(c.Param("id"))
	c.JSON(200, val)
}
