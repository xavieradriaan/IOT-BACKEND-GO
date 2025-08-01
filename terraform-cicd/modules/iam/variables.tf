variable "project_name" {
  description = "Project name"
  type        = string
}

variable "bucket_arn" {
  description = "S3 bucket ARN for artifacts"
  type        = string
}

variable "github_connection_arn" {
  description = "GitHub CodeStar connection ARN"
  type        = string
} 