
terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.14.0"
    }
  }
}

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
