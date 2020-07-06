package catalog

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swagger "github.com/techsimplifier/test-sigma-catalog-services-api/go-client"
)

// RunGetCategories executes the quote read command
func RunGetCategories() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		endpoint := viper.GetString("endpoint")

		config := &swagger.Configuration{
			BasePath:      endpoint,
			DefaultHeader: make(map[string]string),
			UserAgent:     "Swagger-Codegen/1.0.0/go",
		}
		config.AddDefaultHeader("Accept", "application/xml")
		catalogServicesClient := swagger.NewAPIClient(config)
		res, err := catalogServicesClient.CategoriesApi.ApiCategoriesGet(context.Background())
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(body))

	}
}
