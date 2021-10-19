package router

import (
	"github.com/CaiNiaoHui/findkubernetes/router/kubernetes"
	"github.com/CaiNiaoHui/findkubernetes/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Kubernetes kubernetes.RouterGroup
}

var RouterGroupApp = new(RouterGroup)


