package iv_item

import (
	domain "STACK-GO/modules/iv_item/domain"
	intfService "STACK-GO/modules/iv_item/interface"

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
	route.GET("/item/:id/", http.FindByID)
	route.GET("/item/list/", http.List)
	route.POST("/item/add/", http.Create)
	route.PUT("/item/update/:id/", http.Update)
	route.DELETE("/item/delete/:id/", http.Delete)
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
	var request domain.Item
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val := h.port.Create(request)
	c.JSON(200, val)
}

func (h *Handler) Update(c *gin.Context) {
	var request domain.Item
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
