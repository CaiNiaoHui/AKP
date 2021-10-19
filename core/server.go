package core

import (
	"fmt"
	"github.com/CaiNiaoHui/findkubernetes/global"
	"github.com/CaiNiaoHui/findkubernetes/initialize"
	"github.com/CaiNiaoHui/findkubernetes/utils"
	"go.uber.org/zap"
	"time"
)



func RunWindowsServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.JH_CONFIG.System.Addr)

	s := utils.InitServer(address, Router)

	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.JH_LOG.Info("service run success on ", zap.String("address", address))

	fmt.Printf(`
	welcome to use the kuberneters service operator tools
	use the api to get the kubernetes information
	this version :V0.0.1 alpha
	autoRESTful API address :http://IP%s/swagger/index.html
`, address)
	global.JH_LOG.Error(s.ListenAndServe().Error())
}
