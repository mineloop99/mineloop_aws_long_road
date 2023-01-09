package main

import (
	"fmt"

	"github.com/mineloop99/mineloop-aws-long-road/services/ec2"
)

func main() {
	ec2, _ := ec2.RunEC2FreeInstance()
	fmt.Println(ec2)
}
