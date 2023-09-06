package main

import (
	"context"
	"fmt"

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

	image, err := services.CreateImage(client)
	if err != nil {
		fmt.Printf("failed to create EC2 instance, %v", err)
	}
	emiID := image.ImageId
	fmt.Printf("Created EC2 instance with ID: %s\n", *emiID)

	instance, err := services.CreateInstance(emiID, client)
	if err != nil {
		fmt.Printf("failed to create EC2 instance, %v", err)
	}

	instanceId := instance.Instances[0].InstanceId
	fmt.Printf("Created EC2 instance with ID: %s\n", *instanceId)

	instancesResult, err := services.DescribeInstance(client)
	if err != nil {
		fmt.Printf("failed to describe EC2 instance, %v", err)
	}

	for _, reservation := range instancesResult.Reservations {
		for _, i := range reservation.Instances {
			fmt.Printf("Instance ID: %s, State: %s\n", *i.InstanceId, i.State.Name)
		}
	}

	vpcEndpoints, err := services.DescribeVPCEndPoints(client)
	if err != nil {
		fmt.Printf("failed to describe EC2 VpcEndpoints, %v", err)
	}

	for _, vpcEndpoint := range vpcEndpoints.VpcEndpoints {
		fmt.Printf("VPC Endpoint ID: %s\n", *vpcEndpoint.VpcEndpointId)
	}

	monitor, err := services.MonitorIntace(client)
	if err != nil {
		fmt.Printf("failed to monitor instances, %v", err)
	}
	fmt.Println("monitor instances", monitor.InstanceMonitorings)

	reboot, err := services.RebootIntance(client)
	if err != nil {
		fmt.Printf("failed to reboot instances, %v", err)
	}

	fmt.Println("Rebooted instances", reboot.ResultMetadata)

	start, err := services.StartIntace(client)
	if err != nil {
		fmt.Printf("failed to start instances, %v", err)
	}
	fmt.Println("Started instances", start.StartingInstances)

	stop, err := services.StopInstance(client)
	if err != nil {
		fmt.Printf("failed to stop instances, %v", err)
	}
	fmt.Println("Stopped instances", stop.StoppingInstances)

}
