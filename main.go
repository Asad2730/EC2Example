package main

import (
	"context"

	"github.com/Asad2730/EC2Example/services"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	createdInstanceID, err := services.CreateImage(client)
	if err != nil {
		panic("failed to create EC2 instance, %v", err)
	}
}
