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

output "dz1_public_ip" {
  value = aws_instance.dz1.public_ip
}
