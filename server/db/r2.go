package db

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	
	system "usdrive/config"
)

var (
	client *s3.Client
	clientOnce sync.Once
)

func GetR2Client() *s3.Client {
	clientOnce.Do(func() {
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("auto"),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(system.ENV.R2AccessKeyID, system.ENV.R2SecretAccessKey, "")),
		)
		if err != nil {
			log.Fatalf("Failed to load R2 config: %v", err)
		}

		client = s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(system.ENV.R2API)
			o.UsePathStyle = true
		})
	})

	return client
}
