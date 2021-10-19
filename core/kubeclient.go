package core

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func KubeClientSet() (ClientSet *kubernetes.Clientset) {
	// 使用当前上下文环境
	var kubeconfig *string
	// 相对路径
	workDir, _ := os.Getwd()
	workfile := workDir + "/kubeconfig"

	// 测试环境
	kubeconfig = flag.String("kubeconfig", workfile, "absolute path to the kubeconfig file")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// 根据指定的 config 创建一个新的 clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}


