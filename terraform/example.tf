provider "aws" {
  profile = "bob"
  region  = "ap-northeast-1"
}

resource "aws_instance" "example" {
  # ami           = "ami-afb09dc8"
  ami           = "ami-923d12f5"
  instance_type = "t2.micro"

  # https://www.terraform.io/docs/provisioners/index.html
  # - chef
  # - connection
  # - file
  # - local-exec
  # - remote-exec

  provisioner "local-exec" {
    command = "echo ${aws_instance.example.public_ip} > ip_address.txt"
  }

  provisioner "local-exec" {
    when    = "destroy"
    command = "echo ${aws_instance.example.public_ip} > ip_address.txt"
  }
}

# resource "aws_eip" "ip" {
#   instance = "${aws_instance.example.id}"
# }

# resource "aws_instance" "example2" {
#   ami           = "ami-923d12f5"
#   instance_type = "t2.micro"
# }
