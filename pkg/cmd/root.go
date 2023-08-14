package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd/api"

	"github.com/sp-yduck/kubectl-cluster/pkg/kubeconfig"
	"github.com/sp-yduck/kubectl-cluster/pkg/log"
)

type RootOptions struct {
	configFlags *genericclioptions.ConfigFlags

	verbose bool
	debug   bool

	rawConfig api.Config
	args      []string
}

func NewRootOptions() *RootOptions {
	return &RootOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
	}
}

func NewRootCommand() *cobra.Command {
	o := NewRootOptions()
	cmd := &cobra.Command{
		Use:   "kubectl-cluster",
		Short: "kubectl plugin for cluster context control",
		Long: `kubectl plugin for cluster context control
	
		To list all cluster's information in kubeconfig:
			kubectl-cluster
	
		To switch current context by cluster name:
			kubectl-cluster <cluster-name>
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if o.verbose {
				log.InitLogger(zap.InfoLevel)
			}
			if o.debug {
				log.InitLogger(zap.DebugLevel)
			}

			if err := o.complete(cmd, args); err != nil {
				return err
			}
			return o.run()
		},
	}

	cmd.PersistentFlags().BoolVar(&o.debug, "debug", false, "enable DEBUG level log")
	cmd.PersistentFlags().BoolVarP(&o.verbose, "verbose", "v", false, "enabel INFO level log")
	o.configFlags.AddFlags(cmd.Flags())
	return cmd
}

func (o *RootOptions) complete(cmd *cobra.Command, args []string) error {
	zap.S().Debugf("complete called: %v", args)
	o.args = args

	var err error
	o.rawConfig, err = o.configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		zap.S().Errorf("failed to get rawCondig: %v", err)
		return err
	}

	return nil
}

func (o *RootOptions) run() error {
	zap.S().Debugf("run called: %v", o.args)
	args := o.args
	switch len(args) {
	case 0:
		return o.listClusters()
	case 1:
		return o.switchCurrencContextByClusrter(args[0])
	default:
		return errors.New("number of arguments must be one or zero")
	}
}

func (o *RootOptions) listClusters() error {
	zap.S().Debugf("list clusters: %v", o.rawConfig)
	return kubeconfig.ListClusters(o.rawConfig)
}

func (o *RootOptions) switchCurrencContextByClusrter(cluster string) error {
	kconfig := o.configFlags.ToRawKubeConfigLoader().ConfigAccess().GetDefaultFilename()
	zap.S().Debugf("switch current context of '%s' by cluster '%s': %v", kconfig, cluster, o.rawConfig)
	return kubeconfig.SwitchCurrencContextByClusrter(cluster, o.rawConfig, kconfig)
}
