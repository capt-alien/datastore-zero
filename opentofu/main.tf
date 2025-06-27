provider "aws" {
  region = var.aws_region
}

resource "aws_instance" "dz1" {
  ami           = var.ami_id
  instance_type = var.instance_type
  key_name      = "dz1-key"
  tags = {
    Name = "dz1"
  }
}

resource "aws_eip" "dz1_ip" {
  instance = aws_instance.dz1.id
  vpc      = true
}

output "dz1_public_ip" {
  value = aws_eip.dz1_ip.public_ip
}
