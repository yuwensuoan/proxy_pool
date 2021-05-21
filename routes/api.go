package routes

import (
	"github.com/gin-gonic/gin"
	"proxy_pool/app/controllers"
)

// 注册api路由
func RegisterApiRouter(root *gin.Engine)  {
	webRouter := root.Group("/api/")
	proxyRouter := webRouter.Group("/proxy")
	{
		proxyRouter.GET("/list", controllers.ProxyController{}.GetList)
		proxyRouter.GET("/first", controllers.ProxyController{}.GetFirst)
		proxyRouter.POST("/delete", controllers.ProxyController{}.Delete)
	}
	scheduleRouter := webRouter.Group("/schedule")
	{
		scheduleRouter.GET("/stop", controllers.Schedule{}.Stop)
	}
}
