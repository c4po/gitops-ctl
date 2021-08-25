package cmd

import (
	"context"
	"fmt"
	// "github.com/aws/aws-sdk-go-v2/aws"
	// "encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/viper"
	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"log"
	// "os"
)

var codebuildListCmd = &cobra.Command{
	Use:   "list",
	Short: "List codebuild projects",
	Run: func(cmd *cobra.Command, args []string) {

		var github_token = viper.GetString("github_token")
		var vault_addr = viper.GetString("vault_addr")
		fmt.Println(vault_addr, github_token)
		vaultConfig := &api.Config{
			Address: vault_addr,
		}
		vaultClient, err := api.NewClient(vaultConfig)
		if err != nil {
			fmt.Println(err)
			return
		}

		secret, err := vaultClient.Logical().Write(
			"auth/github/login",
			map[string]interface{}{
				"token": github_token,
			})
		if err != nil {
			fmt.Println(err)
			return
		}

		vault_token := secret.Auth.ClientToken

		vaultClient.SetToken(vault_token)

		data, err := vaultClient.Logical().Read("aws_v2/creds/account_id_819784554124")
		if err != nil {
			panic(err)
		}

		log.Println(fmt.Sprintf("List AWS codebuild projects: %d", 10))

		// cfg, err := config.LoadDefaultConfig(context.TODO())
		// if err != nil {
		// 	log.Fatal(err)
		// }

		options := s3.Options{
			Region: "us-east-1",
			Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
				data.Data["access_key"].(string),
				data.Data["secret_key"].(string),
				data.Data["security_token"].(string))),
		}

		// Create an Amazon S3 service client
		// client := s3.NewFromConfig(cfg)
		client := s3.New(options)

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
