package utils

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/mineloop99/mineloop-aws-long-road/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"gopkg.in/yaml.v3"
)

var awsConfig types.AwsConfig = types.AwsConfig{}

var awsCredentials types.AwsCredentials = types.AwsCredentials{}

func GetClientConfig(sessionId string) aws.Config {
	awsConfig := GetAwsConfig()
	awsCred := GetAwsCredentials()

	return aws.Config{
		Region: awsConfig.Region,
		Credentials: credentials.NewStaticCredentialsProvider(
			awsCred.AwsAccessKeyId,
			awsCred.AwsSecretAccessKey,
			sessionId,
		),
	}
}
func GetAwsConfig() types.AwsConfig {
	if awsConfig != (types.AwsConfig{}) {
		return awsConfig
	}
	yfile, err := ioutil.ReadFile("./.aws/config.yaml")
	if err != nil {
		if strings.Contains(err.Error(), "cannot find the file") {

		} else {
			log.Fatal(err)
		}
	}
	err = yaml.Unmarshal(yfile, &awsConfig)
	if err != nil {
		log.Fatal(err)
	}
	return awsConfig
}

func GetAwsCredentials() types.AwsCredentials {
	if awsCredentials != (types.AwsCredentials{}) {
		return awsCredentials
	}
	yfile, err := ioutil.ReadFile("./.aws/credentials.yaml")
	if err != nil {
		if strings.Contains(err.Error(), "cannot find the file") {

		} else {
			log.Fatal(err)
		}
	}
	err = yaml.Unmarshal(yfile, &awsCredentials)
	if err != nil {
		log.Fatal(err)
	}
	return awsCredentials
}
