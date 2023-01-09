package ec2

import (
	"context"
	"fmt"

	utils "github.com/mineloop99/mineloop-aws-long-road/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// EC2CreateInstanceAPI defines the interface for the RunInstances and CreateTags functions.
// We use this interface to test the functions using a mocked service.
type EC2CreateInstanceAPI interface {
	RunInstances(ctx context.Context,
		params *ec2.RunInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)

	CreateTags(ctx context.Context,
		params *ec2.CreateTagsInput,
		optFns ...func(*ec2.Options)) (*ec2.CreateTagsOutput, error)
}

// MakeInstance creates an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a RunInstancesOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to RunInstances.
func MakeInstance(c context.Context, api EC2CreateInstanceAPI, input *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	return api.RunInstances(c, input)
}

// MakeTags creates tags for an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a CreateTagsOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to CreateTags.
func MakeTags(c context.Context, api EC2CreateInstanceAPI, input *ec2.CreateTagsInput) (*ec2.CreateTagsOutput, error) {
	return api.CreateTags(c, input)
}
func RunEC2FreeInstance() (*ec2.RunInstancesOutput, error) {
	ec2Client := ec2.NewFromConfig(utils.GetClientConfig(""))
	ec2InstanceConfig := utils.GetEc2InstanceConfig("ec2_free")
	minMaxCount := int32(1)
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(ec2InstanceConfig.ImageId),
		InstanceType: types.InstanceType(ec2InstanceConfig.VirtualServerType),
		MinCount:     &minMaxCount,
		MaxCount:     &minMaxCount,
	}
	result, err := MakeInstance(context.TODO(), ec2Client, input)
	if err != nil {
		fmt.Println("Got an error creating an instance:")
		fmt.Println(err)
		return nil, err
	}

	tagInput := &ec2.CreateTagsInput{
		Resources: []string{*result.Instances[0].InstanceId},
		Tags: []types.Tag{
			{
				Key:   &ec2InstanceConfig.TagName,
				Value: &ec2InstanceConfig.Name,
			},
		},
	}

	_, err = MakeTags(context.TODO(), ec2Client, tagInput)
	if err != nil {
		fmt.Println("Got an error tagging the instance:")
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Created tagged instance with ID " + *result.Instances[0].InstanceId)
	return result, nil
}
