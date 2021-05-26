package boostrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"proxy_pool/config"
	_ "proxy_pool/config"
	_ "proxy_pool/database"
	"proxy_pool/routes"
	"proxy_pool/app/middleware"
)

var server *gin.Engine

func init()  {
	server = gin.Default()
	server.Use(middleware.LoggerToFile())
	routes.RegisterWebRouter(server)
	routes.RegisterApiRouter(server)
}

func StartServer()  {
	server.Run(fmt.Sprintf(config.CONFIG.App.Addr))
}