package global

import (
	"github.com/CaiNiaoHui/findkubernetes/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

var (
	JH_VP     *viper.Viper
	JH_LOG    *zap.Logger
	JH_CONFIG config.Server
	JH_KUBECLIENT   *kubernetes.Clientset
)
