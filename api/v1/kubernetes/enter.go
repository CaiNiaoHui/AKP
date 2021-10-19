package kubernetes

import "github.com/CaiNiaoHui/findkubernetes/service"

type ApiGroup struct {
	ImageApi
}

var namespaceServer = service.ServiceGroupApp.KubernetesServiceGroup.NamespaceService



