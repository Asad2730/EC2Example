package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func CreateImage(client *ec2.Client) (*ec2.CreateImageOutput, error) {

	// Example: Create an Amazon Machine Image (AMI)
	input := &ec2.CreateImageInput{
		BlockDeviceMappings: []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &types.EbsBlockDevice{
					VolumeSize: aws.Int32(10), // Size in GB
				},
			},
		},
		Description: aws.String("MY AMI"),
		InstanceId:  aws.String("Replace With My instance id"),
		Name:        aws.String("Replace with my AMI NAme"),
	}

	output, err := client.CreateImage(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func CreateInstance(emiId *string, client *ec2.Client) (*ec2.RunInstancesOutput, error) {

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(*emiId),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	}

	output, err := client.RunInstances(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func DescribeInstance(client *ec2.Client) (*ec2.DescribeInstancesOutput, error) {
	input := &ec2.DescribeInstancesInput{}
	output, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func DescribeVPCEndPoints(client *ec2.Client) (*ec2.DescribeVpcEndpointsOutput, error) {

	input := &ec2.DescribeVpcEndpointsInput{}
	output, err := client.DescribeVpcEndpoints(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func MonitorIntace(client *ec2.Client) (*ec2.MonitorInstancesOutput, error) {
	input := &ec2.MonitorInstancesInput{
		InstanceIds: []string{"Replace with your instance ID"},
	}

	output, err := client.MonitorInstances(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, err
}

func StartIntace(client *ec2.Client) (*ec2.StartInstancesOutput, error) {
	input := &ec2.StartInstancesInput{
		InstanceIds: []string{"Replace with your instance ID"},
	}

	output, err := client.StartInstances(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func StopInstance(client *ec2.Client) (*ec2.StopInstancesOutput, error) {
	input := &ec2.StopInstancesInput{
		InstanceIds: []string{"Replace with your instance ID"},
	}
	output, err := client.StopInstances(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	return output, err
}

func RebootIntance(client *ec2.Client) (*ec2.RebootInstancesOutput, error) {
	input := &ec2.RebootInstancesInput{
		InstanceIds: []string{"Replace with your instance ID"},
	}
	output, err := client.RebootInstances(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return output, nil
}
