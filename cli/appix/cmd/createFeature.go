/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"

	pb "appix/api/appix/v1"
)

// createFeatureCmd represents the createFeature command
var createFeatureCmd = &cobra.Command{
	Use:   "feature",
	Short: "Create a new feature",
	Long: `Create a new feature in the system.
Feature is a key-value pair that used to match Hostgroup and Application.
Like GPU, CPU, etc.

Examples:
  appix create feature --name gpu --value v100 
  appix create feature --name gpu --value a100`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewFeaturesClient(conn)

		var req *pb.CreateFeaturesRequest

		if outFile != "" {
			// Generate template YAML file
			feature := &pb.Feature{
				Name:  "feature-name",
				Value: "feature-value",
			}
			features := []*pb.Feature{feature}

			data, err := yaml.Marshal(features)
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

			var features []*pb.Feature
			if err := yaml.Unmarshal(data, &features); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}

			req = &pb.CreateFeaturesRequest{
				Features: features,
			}
		} else {
			// Create from command line flags
			name, _ := cmd.Flags().GetString("name")
			value, _ := cmd.Flags().GetString("value")

			req = &pb.CreateFeaturesRequest{
				Features: []*pb.Feature{
					{
						Name:  name,
						Value: value,
					},
				},
			}
		}

		resp, err := client.CreateFeatures(ctx, req)
		if err != nil {
			log.Fatalf("create feature failed: %v", err)
		}

		if resp != nil {
			fmt.Printf("Code: %d\n", resp.Code)
			fmt.Printf("Message: %s\n", resp.Message)
			fmt.Printf("Action: %s\n", resp.Action)
		}
	},
}

func init() {
	createCmd.AddCommand(createFeatureCmd)
	createFeatureCmd.Flags().String("name", "", "Name of the feature")
	createFeatureCmd.Flags().String("value", "", "Value of the feature")
}
