variable "ssh_allowed_cidr" {}
variable "region" { default = "ap-northeast-1" }
variable "ami_id" { default = "ami-00c408a8b71d5c614" }
variable "key_pair_name" { default = "bob-key" }

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

resource "aws_ecs_cluster" "foo" {
  name = "test-ecs"
}

resource "aws_s3_bucket" "bob_ecs_test_bucket" {
  bucket = "bob-ecs-test-bucket"
  acl    = "private"

  tags = {
    Name        = "ecs_test_bucket"
    Environment = "Dev"
  }
}

resource "aws_instance" "ecs_test_instance" {
  ami                    = var.ami_id
  instance_type          = "t2.micro"
  iam_instance_profile   = "ecsInstanceRole"
  key_name               = var.key_pair_name
  vpc_security_group_ids = [aws_security_group.test_ecs.id]

  user_data = file("scripts/copy-ecs-config-to-s3")
  tags = {
    Name = "Test EC2 instance"
  }
}

resource "aws_ecs_task_definition" "test_ecs" {
  family                = "service"
  container_definitions = file("config/web-task-definition.json")
}
