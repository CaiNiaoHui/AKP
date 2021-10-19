package main

import (
	"github.com/CaiNiaoHui/findkubernetes/core"
	"github.com/CaiNiaoHui/findkubernetes/global"
)

func main() {
	global.JH_VP = core.Viper()      // 初始化Viper
	global.JH_LOG = core.Zap()       // 初始化zap日志库
	global.JH_KUBECLIENT = core.KubeClientSet()

	core.RunWindowsServer()

}
