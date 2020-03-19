variable "ssh_allowed_cidr" {}
variable "region" { default = "ap-northeast-1" }

provider "aws" {
  version = "~> 2.0"
}

terraform {
  backend "s3" {
    bucket = "bobstrange.tfstate"
    key    = "test-ecs.tfstate"
    region = "ap-northeast-1"
  }
}
resource "aws_security_group" "test_ecs" {
  name        = "test_ecs"
  description = "Security group for ECS testing"
  vpc_id      = "vpc-0586736d"

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.ssh_allowed_cidr]
  }

  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Postgres"
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    self        = true
  }

  ingress {
    description = "Redis"
    from_port   = 6379
    to_port     = 6379
    protocol    = "tcp"
    self        = true
  }


  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "Security group for ECS testing"
  }
}
