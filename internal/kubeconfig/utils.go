package kubeconfig

import (
	"log"
	"sort"

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

func GetClusterContextsMap(config api.Config) (map[string][]string, []string) {
	contexts := config.Contexts
	clmap := map[string][]string{}
	keys := []string{}
	for name, ctx := range contexts {
		keys = append(keys, ctx.Cluster)
		if _, ok := clmap[ctx.Cluster]; !ok {
			clmap[ctx.Cluster] = []string{name}
		} else {
			clmap[ctx.Cluster] = append(clmap[ctx.Cluster], name)
		}
	}
	sort.StringSlice(keys).Sort()
	return clmap, keys
}

func Save(config api.Config) error {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	var filename string
	if kubeConfig.ConfigAccess().IsExplicitFile() {
		filename = kubeConfig.ConfigAccess().GetExplicitFile()
	} else {
		filename = kubeConfig.ConfigAccess().GetDefaultFilename()
	}
	err := clientcmd.WriteToFile(config, filename)
	return err
}
