package kubeconfig

import (
	"log"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetRawConfig() (config api.Config) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.RawConfig()
	if err != nil {
		log.Fatal("couldn't get kubeconfig")
	}
	return config
}

func ReadCurrentCluster(config api.Config) (currentCluster string) {
	contexts := config.Contexts
	currentContext := contexts[config.CurrentContext]
	currentCluster = currentContext.Cluster
	return currentCluster
}
