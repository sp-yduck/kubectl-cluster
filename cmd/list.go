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
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/sp-yduck/kubectl-cluster/internal/kubeconfig"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "view all clusters from your KUBECONFIG",
	Long:  `view all clusters from your KUBECONFIG`,
	RunE:  List,
}

func List(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("current command doesn't accept any subcommands/args")
	}
	config := kubeconfig.GetRawConfig()
	currentCluster := kubeconfig.ReadCurrentCluster(config)
	clmap, clusterNames := kubeconfig.GetClusterContextsMap(config)
	clusters := config.Clusters
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"cluster", "apiserver endpoint", "context."})
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

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
