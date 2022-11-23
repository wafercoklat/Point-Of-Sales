package routes

import (
	"STACK-GO/adapter"

	Auth "STACK-GO/middleware"
	UserHandler "STACK-GO/modules/dt_user/handler"
	UserService "STACK-GO/modules/dt_user/service"
	LoginHandler "STACK-GO/modules/io_loginregister/handler"
	LoginService "STACK-GO/modules/io_loginregister/service"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(g *gin.Engine, db adapter.RepoAdapter) {
	// Init EndPoint
	login := LoginHandler.NewHandler(LoginService.NewService(db))
	user := UserHandler.NewHandler(UserService.NewService(db))

	// EndPoint
	// Public
	public := g.Group("/")
	public.POST("/login", login.LoginCheck)

	// Middleware
	private := g.Group("/")
	private.Use(Auth.JwtAuthMiddleware())

	// User
	private.GET("/user/:id/", user.FindByID)
	private.GET("/user/list/", user.List)
	private.POST("/user/add/", user.Create)
	private.PUT("/user/update/:id/", user.Update)
	private.DELETE("/user/delete/:id/", user.Delete)
}
