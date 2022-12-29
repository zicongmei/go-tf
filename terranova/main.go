package main

import (
	"fmt"

	"github.com/johandry/terranova"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

var code = `
provider "aws" {
  region = "us-east-1"
}

resource "aws_vpc" "this" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
	Name = "zicong-test-vpc"
  }
}  
`

const stateFilename = "terraform.tfstate"

func main() {
	platform := terranova.NewPlatform(code)
	fmt.Println(platform.Code)

	platform.AddProvider("aws", aws.Provider())
	fmt.Println("added aws provider")
}
