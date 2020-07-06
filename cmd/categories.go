package cmd

import (
	"github.com/spf13/cobra"
	"github.com/techsimplifier/test-sigma-catalog-services-api/pkg/catalog"
)

// categoriesCmd represents the categories command
var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Calls GET /api/categories",
	Long:  `Calls GET /api/categories.`,
	Run:   catalog.RunGetCategories(),
}

func init() {
	catalogCmd.AddCommand(categoriesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoriesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoriesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
