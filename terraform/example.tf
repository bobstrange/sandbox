provider "aws" {
  profile = "bob"
  region  = "ap-northeast-1"
}

resource "aws_instance" "example" {
  # ami           = "ami-afb09dc8"
  ami           = "ami-923d12f5"
  instance_type = "t2.micro"
}

resource "aws_eip" "ip" {
  instance = "${aws_instance.example.id}"
}

resource "aws_instance" "example2" {
  # ami           = "ami-afb09dc8"
  ami           = "ami-923d12f5"
  instance_type = "t2.micro"
}
