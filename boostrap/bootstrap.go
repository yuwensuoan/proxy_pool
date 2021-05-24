package boostrap

import (
	"github.com/gin-gonic/gin"
	_ "proxy_pool/config"
	_ "proxy_pool/database"
	"proxy_pool/routes"
	"proxy_pool/app/middleware"
)

var Server *gin.Engine

func init()  {
	Server = gin.Default()
	Server.Use(middleware.LoggerToFile())
	routes.RegisterWebRouter(Server)
	routes.RegisterApiRouter(Server)
}