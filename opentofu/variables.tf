variable "aws_region" {
  description = "The AWS region to deploy to"
  type        = string
  default     = "us-west-2"
}
variable "ami_id" {
  description = "The AMI to use for the instance"
}
variable "instance_type" {
  default = "t3.small"
}
