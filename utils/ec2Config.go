package utils

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/mineloop99/mineloop-aws-long-road/types"

	"gopkg.in/yaml.v3"
)

func GetEc2InstanceConfig(ec2Name string) types.Ec2Instance {
	yfile, err := ioutil.ReadFile("./.aws/ec2/" + ec2Name + ".yaml")
	if err != nil {
		if strings.Contains(err.Error(), "cannot find the file") {

		} else {
			log.Fatal(err)
		}
	}
	ec2Instance := types.Ec2Instance{}
	err = yaml.Unmarshal(yfile, &ec2Instance)
	if err != nil {
		log.Fatal(err)
	}
	return ec2Instance
}
