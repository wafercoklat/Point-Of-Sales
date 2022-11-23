package fi_receiptvoucher

import (
	respondErr "STACK-GO/error"
	domain "STACK-GO/modules/fi_receiptvoucher/domain"
	intfService "STACK-GO/modules/fi_receiptvoucher/intrface"

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
	route.GET("/receiptvoucher/:id/", http.FindByID)
	route.GET("/receiptvoucher/list/", http.List)
	route.POST("/receiptvoucher/add/", http.Create)
	route.PUT("/receiptvoucher/update/:id/", http.Update)
	route.DELETE("/receiptvoucher/delete/:id/", http.Delete)
}

func (h *Handler) FindByID(c *gin.Context) {
	val, err := h.port.FindByID(c.Param("id"))
	if respondErr.Error(c, err) {
		return
	}
	c.JSON(200, val)
}

func (h *Handler) List(c *gin.Context) {
	val, err := h.port.List()
	if respondErr.Error(c, err) {
		return
	}
	c.JSON(200, val)
}

func (h *Handler) Create(c *gin.Context) {
	var request domain.ReceiptvoucherPackage
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val, err := h.port.Create(domain.ReceiptvoucherPackage{})
	if respondErr.Error(c, err) {
		return
	}
	c.JSON(200, val)
}

func (h *Handler) Update(c *gin.Context) {
	var request domain.ReceiptvoucherPackage
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val, err := h.port.Update(c.Param("id"), request)
	if respondErr.Error(c, err) {
		return
	}
	c.JSON(200, val)
}

func (h *Handler) Delete(c *gin.Context) {
	val, err := h.port.Delete(c.Param("id"))
	if respondErr.Error(c, err) {
		return
	}
	c.JSON(200, val)
}
