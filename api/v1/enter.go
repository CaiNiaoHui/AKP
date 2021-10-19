package v1

import (
	"github.com/CaiNiaoHui/findkubernetes/api/v1/kubernetes"
	"github.com/CaiNiaoHui/findkubernetes/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	KubernetesApiGroup kubernetes.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

