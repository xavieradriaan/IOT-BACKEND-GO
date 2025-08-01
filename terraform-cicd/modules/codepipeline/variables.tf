variable "project_name" {
  description = "Project name"
  type        = string
}

variable "service_role_arn" {
  description = "IAM role ARN for CodePipeline"
  type        = string
}

variable "artifacts_bucket" {
  description = "S3 bucket for artifacts"
  type        = string
}

variable "github_repo" {
  description = "GitHub repository (owner/repo)"
  type        = string
}

variable "github_branch" {
  description = "GitHub branch"
  type        = string
}

variable "github_connection_arn" {
  description = "GitHub connection ARN"
  type        = string
}

variable "codebuild_project" {
  description = "CodeBuild project name"
  type        = string
}

variable "codedeploy_app" {
  description = "CodeDeploy application name"
  type        = string
}

variable "deployment_group" {
  description = "CodeDeploy deployment group name"
  type        = string
} 