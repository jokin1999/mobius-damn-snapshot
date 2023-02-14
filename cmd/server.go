package cmd

import (
	"github.com/jokin1999/mobius-damn-snapshot/services/bunny"
	"github.com/spf13/cobra"
)

func init() {
	var serverPort = ""

	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start a mobius server.",
		Run: func(cmd *cobra.Command, args []string) {
			bunny.Init()
		},
	}

	// specify port
	// This flag has a higher priority than configuration and environmnet variables.
	serverCmd.Flags().StringVarP(&serverPort, "port", "p", "", "specify server listening tcp port")

	rootCmd.AddCommand(serverCmd)
}
