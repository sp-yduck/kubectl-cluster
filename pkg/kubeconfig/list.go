/*
Copyright © 2023 Teppei Sudo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package kubeconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
)

func ListClusters() error {
	config, err := getRawConfig()
	if err != nil {
		zap.S().Errorf("failed to get kubeconfig: %v", err)
		return err
	}
	currentCluster, err := readCurrentCluster(*config)
	if err != nil {
		zap.S().Errorf("failed to read current cluster from kubeconfig: %v", err)
		return err
	}
	clmap, clusterNames := getClusterContextsMap(*config)
	clusters := config.Clusters
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"cluster", "apiserver endpoint", "context"})
	for _, name := range clusterNames {
		var clname string
		if name == currentCluster {
			clname = fmt.Sprintf("-> %s", name)
		} else {
			clname = fmt.Sprintf("   %s", name)
		}
		table.Append([]string{clname, clusters[name].Server, strings.Join(clmap[name], "\n")})
	}
	table.Render()
	return nil
}