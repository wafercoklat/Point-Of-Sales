package modules

import (
	"STACK-GO/adapter"
	UserHandler "STACK-GO/modules/user/handler"
	UserService "STACK-GO/modules/user/service"

	"github.com/gin-gonic/gin"
)

func New(g *gin.Engine, db adapter.RepoAdapter) {
	UserHandler.NewHandler(g, UserService.NewService(db))
}
