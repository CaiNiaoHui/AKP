package kubernetes

import (
	"fmt"
	"github.com/CaiNiaoHui/findkubernetes/global"
	"github.com/CaiNiaoHui/findkubernetes/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

type ImageApi struct {
}


// @Tags Kubernetes
// @Summary getAllImage
// @Schemes
// @Description from the kubernetes get all images
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"success"}"
// @Router /kubernetes/image/getAllImageList [get]
func (j *ImageApi) GetAllImageList(c *gin.Context) {
	// 获取 namespaces
	nsList, err := global.JH_KUBECLIENT.CoreV1().Namespaces().List(c, metav1.ListOptions{})
	if err != nil {
		global.JH_LOG.Error("no namespace source data", zap.Any("err", err))
		response.Fail(c, nil, "don't find source")
		return
	}

	// -- 从这里开始使用协程   已经拿到数组 nsList.Items
	// 1.

	var ImageArr []string
	for i := 0; i < len(nsList.Items); i++ {
		// find pod list
		podList, err := global.JH_KUBECLIENT.CoreV1().Pods(nsList.Items[i].Name).List(c, metav1.ListOptions{})
		if err != nil {
			global.JH_LOG.Error("no list Pod source data", zap.Any("err", err))
			response.Fail(c, nil, "don't find source")
			return
		}
		// get each pod information
		for j := 0; j < len(podList.Items); j++ {
			podName := podList.Items[j].Name
			eachPod , err := global.JH_KUBECLIENT.CoreV1().Pods(nsList.Items[i].Name).Get(c, podName, metav1.GetOptions{})
			if err != nil {
				global.JH_LOG.Error("no get Pod source data", zap.Any("err", err))
				response.Fail(c, nil, "don't find source")
				return
			}
			// get each containers information
			for k := 0; k < len(eachPod.Spec.Containers); k++ {
				eachImage := eachPod.Spec.Containers[k].Image
				ImageArr = append(ImageArr, eachImage)
			}
		}
	}
	response.Success(c, gin.H{"eachImageArr": ImageArr}, "success")
	fmt.Println(ImageArr)

}

