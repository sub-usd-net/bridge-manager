package cmd

import "github.com/spf13/cobra"

func New() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "bridge",
		Short: "C-Chain <-> Subnet Bridge Manager Services and Utilities",
	}

	rootCmd.AddCommand(serviceCmd())

	return rootCmd, nil
}
