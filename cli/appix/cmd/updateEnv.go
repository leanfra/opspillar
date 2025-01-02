/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "appix/api/appix/v1"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/yaml.v2"
)

// updateEnvCmd represents the updateEnv command
var updateEnvCmd = &cobra.Command{
	Use:     "env",
	Short:   "Update an environment",
	Long:    `Update an environment with the specified ID and fields.`,
	Aliases: []string{"env", "envs", "environment"},
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to connect: %v\n", err)
			return
		}
		defer conn.Close()
		client := pb.NewEnvsClient(conn)

		var envs []*pb.Env
		if updateOnline {
			// Get existing environment data
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for online editing")
			}

			// Get the environment data
			getReq := &pb.GetEnvsRequest{Id: id}
			getResp, err := client.GetEnvs(cmd.Context(), getReq)
			if err != nil {
				log.Fatalf("failed to get environment: %v", err)
			}

			// Convert to YAML
			data, err := yaml.Marshal([]*pb.Env{getResp.Env})
			if err != nil {
				log.Fatalf("failed to marshal environment: %v", err)
			}

			// Create temp file
			tmpfile, err := os.CreateTemp("", "env-*.yaml")
			if err != nil {
				log.Fatalf("failed to create temp file: %v", err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write(data); err != nil {
				log.Fatalf("failed to write temp file: %v", err)
			}
			tmpfile.Close()

			editor := findEditor()
			if editor == "" {
				log.Fatal("no suitable editor found - please set EDITOR environment variable")
			}

			cmd := exec.Command(editor, tmpfile.Name())
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatalf("failed to run editor: %v", err)
			}

			// Read updated content
			updatedData, err := os.ReadFile(tmpfile.Name())
			if err != nil {
				log.Fatalf("failed to read updated file: %v", err)
			}

			// Parse updated YAML
			if err := yaml.Unmarshal(updatedData, &envs); err != nil {
				log.Fatalf("failed to parse updated yaml: %v", err)
			}

		} else if updateFile != "" {
			// Read and parse YAML file
			data, err := os.ReadFile(updateFile)
			if err != nil {
				log.Fatalf("failed to read yaml file: %v", err)
			}

			if err := yaml.Unmarshal(data, &envs); err != nil {
				log.Fatalf("failed to parse yaml: %v", err)
			}
		} else {
			// Command line update
			id, _ := cmd.Flags().GetUint32("id")
			if id == 0 {
				log.Fatal("id is required for command line update")
			}
			name, _ := cmd.Flags().GetString("name")
			description, _ := cmd.Flags().GetString("description")

			envs = []*pb.Env{
				{
					Id:          id,
					Name:        name,
					Description: description,
				},
			}
		}

		req := &pb.UpdateEnvsRequest{
			Envs: envs,
		}

		reply, err := client.UpdateEnvs(cmd.Context(), req)
		if err != nil {
			fmt.Printf("Error updating environment: %v\n", err)
			return
		}

		if reply != nil {
			fmt.Printf("Action: %s\n", reply.Action)
			fmt.Printf("Code: %d\n", reply.Code)
			fmt.Printf("Message: %s\n", reply.Message)
		}
	},
}

func init() {
	updateCmd.AddCommand(updateEnvCmd)

	updateEnvCmd.Flags().Uint32("id", 0, "Environment ID to update")
	updateEnvCmd.Flags().String("name", "", "New environment name")
	updateEnvCmd.Flags().String("desc", "", "New environment description")

}
