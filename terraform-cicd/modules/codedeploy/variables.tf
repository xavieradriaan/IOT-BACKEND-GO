variable "project_name" {
  description = "Project name"
  type        = string
}

variable "service_role_arn" {
  description = "IAM role ARN for CodeDeploy"
  type        = string
}

variable "ec2_tag_key" {
  description = "EC2 tag key for deployment"
  type        = string
}

variable "ec2_tag_value" {
  description = "EC2 tag value for deployment"
  type        = string
} 