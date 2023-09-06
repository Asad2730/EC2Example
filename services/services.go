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
