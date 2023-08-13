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
	"log"

	"github.com/spf13/cobra"

	"github.com/sp-yduck/kubectl-cluster/pkg/kubeconfig"
)

func use(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("switch command accepts only one cluster name")
	}
	config, err := kubeconfig.GetRawConfig()
	if err != nil {
		return err
	}
	clusters := map[string]string{}
	for name, context := range config.Contexts {
		clusters[context.Cluster] = name
	}
	contextName, ok := clusters[args[0]]
	if !ok {
		log.Fatalf("there is no context using cluster \"%s\"\n", args[0])
	}
	config.CurrentContext = contextName
	if err := kubeconfig.Save(*config); err != nil {
		return err
	}
	fmt.Printf("switched to cluster \"%s\" (context: \"%s\")\n", args[0], contextName)
	return nil
}
