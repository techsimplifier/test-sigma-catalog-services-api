package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// servicesCmd represents the services command
var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Access services exposed by the Sigma Catalog Services API",
	Long:  `Access services exposed by the Sigma Catalog Services API.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	rootCmd.AddCommand(catalogCmd)

	catalogCmd.PersistentFlags().StringP("endpoint", "", "https://localhost", "Endpoint for the Sigma Catalog Services swagger API")

	// Bind flags from the command line to the viper framework
	if err := viper.BindPFlags(catalogCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}
