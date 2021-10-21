package kubernetes

import (
	"fmt"
	"github.com/CaiNiaoHui/findkubernetes/global"
	"github.com/CaiNiaoHui/findkubernetes/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"runtime"
	"sync"
	"time"
)

type ImageApi struct {
}

var foundMatch = make(chan string)	// 传输结果
var ImageArr []string	// 记录结果
var workerCount = 0

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
	// 0. 计时
	start := time.Now()

	// 1. 使用协程锁
	var wg sync.WaitGroup
	maxGoroutines := runtime.NumCPU() * 8
	limiter := make(chan struct{}, maxGoroutines)	//chan struct 内存占用0 bool占用1

	for i := 0; i < len(nsList.Items); i++ {
		// find ns Name and pod list
		nsName := nsList.Items[i].Name
		podList, err := global.JH_KUBECLIENT.CoreV1().Pods(nsName).List(c, metav1.ListOptions{})
		if err != nil {
			global.JH_LOG.Error("no list Pod source data", zap.Any("err", err))
			response.Fail(c, nil, "don't find source")
			return
		}

		wg.Add(1)
		limiter <- struct{}{}

		go func(*v1.PodList, string, []string) {
			defer func() {
				wg.Done()
				<- limiter
			}()
			// get each pod information
			for j := 0; j < len(podList.Items); j++ {
				podName := podList.Items[j].Name
				eachPod , err := global.JH_KUBECLIENT.CoreV1().Pods(nsName).Get(c, podName, metav1.GetOptions{})
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
		}(podList, nsName, ImageArr)

	}
	wg.Wait()
	response.Success(c, gin.H{"eachImageArr": ImageArr}, "success")
	//fmt.Println(ImageArr)
	fmt.Println(time.Since(start))
}


