package cmd

import (
	"context"
	"fmt"
	// "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"encoding/json"
	"log"
	"os"
)

var codebuildListCmd = &cobra.Command{
	Use:   "list",
	Short: "List codebuild projects",
	Run: func(cmd *cobra.Command, args []string) {

		var token = os.Getenv("VAULT_TOKEN")
		var vault_addr = os.Getenv("VAULT_ADDR")

		vaultConfig := &api.Config{
			Address: vault_addr,
		}
		vaultClient, err := api.NewClient(vaultConfig)
		if err != nil {
			fmt.Println(err)
			return
		}
		vaultClient.SetToken(token)

		data, err := vaultClient.Logical().Read("aws_v2/creds/account_id_819784554124")
		if err != nil {
			panic(err)
		}

		b, _ := json.Marshal(data.Data)
		fmt.Println(string(b))

		log.Println(fmt.Sprintf("List AWS codebuild projects: %d", 10))

		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		// Create an Amazon S3 service client
		client := s3.NewFromConfig(cfg)

		input := &s3.ListBucketsInput{}

		result, err := client.ListBuckets(context.TODO(), input)
		if err != nil {
			fmt.Println("Got an error retrieving buckets:")
			fmt.Println(err)
			return
		}

		fmt.Println("Buckets:")

		for _, bucket := range result.Buckets {
			fmt.Println(*bucket.Name + ": " + bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
		}

	},
}

func init() {
	codebuildCmd.AddCommand(codebuildListCmd)
}
