package service

import (
	"github.com/CaiNiaoHui/findkubernetes/service/kubernetes"
)

type ServiceGroup struct {
	KubernetesServiceGroup kubernetes.ServiceGroup

}

var ServiceGroupApp = new(ServiceGroup)
