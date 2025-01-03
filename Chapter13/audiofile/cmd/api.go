package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter13/audiofile/utils"

	metadataService "github.com/ibiscum/Building-Modern-CLI-Applications-in-Go/Chapter13/audiofile/services/metadata"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start or stop the API required by the CLI",
	Long: `Start or stop the API with the following usage:
./audiofile api`,
	Example: `audiofile api`,
	Run: func(cmd *cobra.Command, args []string) {
		configure()
		utils.InitAPILogger()
		var port int
		flag.IntVar(&port, "p", viper.GetInt("api.port"), "Port for metadata service")
		flag.Parse()
		fmt.Printf("Starting API at http://localhost:%d\nPress Ctrl-C to stop.\n", port)
		metadataService.Run(port)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func configure() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("api")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetDefault("api.port", 8000)
}
