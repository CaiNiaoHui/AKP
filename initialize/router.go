package initialize

import (
	_ "github.com/CaiNiaoHui/findkubernetes/docs"
	"github.com/CaiNiaoHui/findkubernetes/global"
	"github.com/CaiNiaoHui/findkubernetes/middleware"
	"github.com/CaiNiaoHui/findkubernetes/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// cor
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.JH_LOG.Info("use middleware cors")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.JH_LOG.Info("register swagger handler")


	kubernetesRouter := router.RouterGroupApp.Kubernetes

	PublicGroup := Router.Group("kubernetes")
	{
		kubernetesRouter.InitImageRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}


	global.JH_LOG.Info("router register success")
	//Router.Run(":"+ global.JH_CONFIG.System.Addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	return Router
}



