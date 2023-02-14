package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const banner = `
███╗   ███╗ ██████╗ ██████╗ ██╗██╗   ██╗███████╗
████╗ ████║██╔═══██╗██╔══██╗██║██║   ██║██╔════╝
██╔████╔██║██║   ██║██████╔╝██║██║   ██║███████╗
██║╚██╔╝██║██║   ██║██╔══██╗██║██║   ██║╚════██║
██║ ╚═╝ ██║╚██████╔╝██████╔╝██║╚██████╔╝███████║
╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚═╝ ╚═════╝ ╚══════╝

DAMN-Snapshot-Controller
`

var rootCmd = &cobra.Command{
	Use:   "mobius-damn-snapshot-controller",
	Short: "Mobius-damn-snapshot-controller is a fast controller for pve snapshots.",
	Run: func(cmd *cobra.Command, args []string) {
		// print banner
		fmt.Print(banner)

		// show help information
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
