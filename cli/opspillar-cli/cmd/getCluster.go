package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	pb "opspillar/api/opspillar/v1"
)

var getClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Get clusters",
	Long: `Get clusters resources from the system.

Examples:
  opspillar get cluster                              # List all
  opspillar get cluster --names cluster1,cluster2    # Filter by names
  opspillar get cluster --ids 1,2,3                 # Filter by IDs
  opspillar get cluster --page 1 --page-size 10     # With pagination
  opspillar get cluster --names prod --format yaml   # Combined filters`,
	Aliases: []string{"clusters", "cls"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, conn, err := NewConnection(true)
		if err != nil {
			log.Fatalf("connect to server failed: %v", err)
		}
		defer conn.Close()

		client := pb.NewClustersClient(conn)

		page := GetPage
		pageSize := GetPageSize
		names, _ := cmd.Flags().GetStringSlice("names")
		uintIds, _ := cmd.Flags().GetUintSlice("ids")
		ids := make([]uint32, len(uintIds))
		for i, id := range uintIds {
			ids[i] = uint32(id)
		}

		var allClusters []*pb.Cluster
		for {
			req := &pb.ListClustersRequest{
				Page:     page,
				PageSize: pageSize,
				Names:    names,
				Ids:      ids,
			}

			resp, err := client.ListClusters(ctx, req)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			if resp.Code != 0 {
				fmt.Printf("Response details:\n")
				fmt.Printf("  Message: %s\n", resp.Message)
				fmt.Printf("  Code: %d\n", resp.Code)
				fmt.Printf("  Action: %s\n", resp.Action)
				return
			}

			// 添加当前页的集群到结果集
			allClusters = append(allClusters, resp.Clusters...)

			// 如果返回的集群数量小于页大小，说明已经是最后一页
			if len(resp.Clusters) < int(pageSize) {
				break
			}

			page++
		}

		switch GetFormat {
		case "yaml":
			data, err := yaml.Marshal(allClusters)
			if err != nil {
				log.Fatalf("序列化YAML失败: %v", err)
			}
			fmt.Println(string(data))
		case "table":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"ID", "Name", "Description"})
			for _, cluster := range allClusters {
				table.Append([]string{
					fmt.Sprintf("%d", cluster.Id),
					cluster.Name,
					cluster.Description,
				})
			}
			table.Render()
		case "text":
			if len(allClusters) == 0 {
				fmt.Println("No clusters found")
				return
			}
			for _, cluster := range allClusters {
				fmt.Printf("ID: %d \t Name: %s \t Description: %s\n",
					cluster.Id, cluster.Name, cluster.Description)
			}
		default:
			fmt.Println("unknown format")
		}
	},
}

func init() {
	getCmd.AddCommand(getClusterCmd)

	getClusterCmd.Flags().StringSlice("names", []string{}, "Filter by cluster names")
	getClusterCmd.Flags().UintSlice("ids", []uint{}, "Filter by cluster IDs")
}
