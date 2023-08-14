package kubeconfig

import (
	"errors"
	"fmt"
	"sort"

	"k8s.io/client-go/tools/clientcmd/api"
)

func readCurrentCluster(config api.Config) (string, error) {
	if config.CurrentContext == "" {
		return "", errors.New("current context is not present")
	}
	contexts := config.Contexts
	currentContext, ok := contexts[config.CurrentContext]
	if !ok {
		return "", fmt.Errorf("current context '%s' is not found in contexts", config.CurrentContext)
	}
	if currentContext.Cluster == "" {
		return "", fmt.Errorf("context '%s' does not have cluster info", config.CurrentContext)
	}
	return currentContext.Cluster, nil
}

func getClusterContextsMap(config api.Config) (map[string][]string, []string) {
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
