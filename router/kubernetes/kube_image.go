package kubernetes

import (
	v1 "github.com/CaiNiaoHui/findkubernetes/api/v1"
	"github.com/gin-gonic/gin"
)

type ImageRouter struct {
	
}

func (s *ImageRouter) InitImageRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	imageRouter := Router.Group("image")
	var ImageApi = v1.ApiGroupApp.KubernetesApiGroup.ImageApi
	{
		imageRouter.GET("getAllImageList", ImageApi.GetAllImageList)
	}



	return imageRouter
}

