package kubeconfig

import (
	"errors"
	"fmt"
	"sort"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func GetRawConfig() (*api.Config, error ) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func ReadCurrentCluster(config api.Config) (string, error) {
	if config.CurrentContext == "" {
		return "", errors.New("current context is not present")
	}
	contexts := config.Contexts
	currentContext, ok := contexts[config.CurrentContext]
	if !ok {
		return "", fmt.Errorf("current context %s is not found in contexts", config.CurrentContext)
	}
	return currentContext.Cluster, nil
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
