variable "aws_region" {
  description = "The AWS region to deploy to"
  type        = string
  default     = "eu-north-1"
}
variable "ami_id" {
  description = "The AMI to use for the instance"
}
variable "instance_type" {
  default = "t3.micro"
}
