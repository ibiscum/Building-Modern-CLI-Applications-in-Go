package cmd

import (
	"fmt"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter09/audiofile/utils"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Version of audiofile CLI",
	Example: `audiofile version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
