package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sub-usd-net/bridge-manager/pkg/service"
	"github.com/sub-usd-net/bridge-manager/pkg/types"
	"go.uber.org/zap"
)

func serviceCmd() *cobra.Command {
	configFile := ""
	command := &cobra.Command{
		Use:   "service",
		Short: "Bridging Service",
	}
	command.PersistentFlags().StringVar(&configFile, "config", "config.yaml", "path to config file")

	command.RunE = func(_ *cobra.Command, _ []string) error {
		cfg, err := types.NewBridgeConfigFromPath(configFile)
		if err != nil {
			return err
		}

		log, err := zap.NewProduction()
		if err != nil {
			return err
		}

		svc, err := service.NewBridgeService(cfg, log.Sugar())
		if err != nil {
			return err
		}

		return svc.Start()
	}
	return command
}
