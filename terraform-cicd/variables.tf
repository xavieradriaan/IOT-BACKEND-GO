variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "project_name" {
  description = "Project name"
  type        = string
  default     = "iot-backend"
}

variable "github_owner" {
  description = "GitHub account or organization"
  type        = string
}

variable "github_repo" {
  description = "GitHub repository name (only)"
  type        = string
}

variable "github_branch" {
  description = "Branch to track for CI/CD"
  type        = string
  default     = "main"
}

variable "github_connection_arn" {
  description = "GitHub CodeStar connection ARN"
  type        = string
}

variable "ec2_tag_key" {
  description = "Tag key of EC2 instance used in CodeDeploy"
  type        = string
  default     = "Name"
}

variable "ec2_tag_value" {
  description = "Tag value of EC2 instance (must match an instance running with CodeDeploy agent)"
  type        = string
  default     = "iot-backend-server"
}
