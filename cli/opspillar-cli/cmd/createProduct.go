/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	pb "opspillar/api/opspillar/v1"
)

// createProductCmd represents the createProduct command
var createProductCmd = &cobra.Command{
	Use:   "product",
	Short: "Create a new product",
	Long: `Create a new product in the system.
Product is a business unit that used for grouping resources.

Examples:
  opspillar create product --name web-app --code webapp --desc "Web Application"
  opspillar create product --name mobile-app --code mobile --desc "Mobile Application"`,
	Aliases: []string{"prod", "prods", "products"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewProductsClient(conn)

		var req *pb.CreateProductsRequest

		if outFile != "" {
			// Generate template YAML file
			product := &pb.Product{
				Name:        "product-name",
				Code:        "product-code",
				Description: "product description",
			}
			products := []*pb.Product{product}

			data, err := yaml.Marshal(products)
			if err != nil {
				log.Fatalf("failed to generate yaml: %v", err)
			}

			if err := os.WriteFile(outFile, data, 0644); err != nil {
				log.Fatalf("failed to write template file: %v", err)
			}

			fmt.Printf("Template file generated at: %s\n", outFile)
			return
		} else if yamlFile != "" {
			// Read from YAML file
			data, err := os.ReadFile(yamlFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			var products []*pb.Product
			if err := yaml.Unmarshal(data, &products); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateProductsRequest{
				Products: products,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			code, _ := cmd.Flags().GetString("code")
			desc, _ := cmd.Flags().GetString("desc")

			req = &pb.CreateProductsRequest{
				Products: []*pb.Product{
					{
						Name:        name,
						Code:        code,
						Description: desc,
					},
				},
			}
		}

		resp, err := client.CreateProducts(ctx, req)
		if err != nil {
			log.Fatalf("failed to create products: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createProductCmd)
	createProductCmd.Flags().String("name", "", "Name of the product")
	createProductCmd.Flags().String("code", "", "Code of the product")
	createProductCmd.Flags().String("desc", "", "Description of the product")
}
