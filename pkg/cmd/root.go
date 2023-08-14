package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/sp-yduck/kubectl-cluster/pkg/log"
)

type RootOptions struct {
	configFlags *genericclioptions.ConfigFlags

	rawConfig      api.Config
	args           []string
}

func NewRootOptions() RootOptions {
	return RootOptions{}
}


func (o *RootOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args

	var err error
	o.rawConfig, err = o.configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return err
	}

	return nil
}


func (o *RootOptions)run() error {
	args := o.args
	zap.S().Debugf("root command called: args=%v", args)
	switch len(args) {
	case 0:
		return kubeconfig.ListClusters()
	case 1:
		return kubeconfig.SwitchCurrencContextByClusrter(args[0])
	default:
		return errors.New("number of arguments must be one or zero")
	}
}	

func NewRootCommand() *cobra.Command {
	o := NewRootOptions()

	cmd := &cobra.Command{
		Use:   "kubectl-cluster",
		Short: "kubectl plugin for cluster context control",
		Long:  `kubectl plugin for cluster context control
	
		To list all cluster's information in kubeconfig:
			kubectl-cluster
	
		To switch current context by cluster name:
			kubectl-cluster <cluster-name>
		`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				log.InitLogger(zap.InfoLevel)
			} 
			if debug {
				log.InitLogger(zap.DebugLevel)
			}
			
			if err := o.Complete(cmd, args); err != nil {
				return err
			}
			return o.run()
		},
	}

	cmd.Flags().BoolVar(&o.listNamespaces, "list", o.listNamespaces, "if true, print the list of all namespaces in the current KUBECONFIG")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}