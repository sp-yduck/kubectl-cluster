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
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/sp-yduck/kubectl-cluster/pkg/log"	
)

var verbose bool
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-cluster",
	Short: "kubectl plugin for cluster context control",
	Long:  `kubectl plugin for cluster context control`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			log.InitLogger(zap.InfoLevel)
		} 
		if debug {
			log.InitLogger(zap.DebugLevel)
		}
		return runRoot(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runRoot(cmd *cobra.Command, args []string) error {
	zap.S().Debugf("root command called: args=%v", args)
	if len(args) == 0 {
		return list(cmd)
	}
	return use(cmd, args)
}

func init() {
	cobra.OnInitialize(func() {
		log.InitLogger(zap.FatalLevel)
	})
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "kubeconfig", "", "config file (default is $HOME/.kube/config)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,  "enable info level log")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug level log")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
