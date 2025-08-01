terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# Data sources
data "aws_caller_identity" "current" {}

# S3 bucket for artifacts  
resource "aws_s3_bucket" "codepipeline_artifacts" {
  bucket        = "${var.project_name}-pipeline-artifacts-${random_string.bucket_suffix.result}"
  force_destroy = true
}

resource "random_string" "bucket_suffix" {
  length  = 8
  special = false
  upper   = false
}

resource "aws_s3_bucket_versioning" "codepipeline_artifacts" {
  bucket = aws_s3_bucket.codepipeline_artifacts.id
  versioning_configuration {
    status = "Enabled"
  }
}

# Modules
module "iam_roles" {
  source               = "./modules/iam"
  project_name         = var.project_name
  bucket_arn           = aws_s3_bucket.codepipeline_artifacts.arn
  github_connection_arn = var.github_connection_arn
}

module "codebuild" {
  source           = "./modules/codebuild"
  project_name     = var.project_name
  service_role_arn = module.iam_roles.codebuild_role_arn
}

module "codedeploy" {
  source           = "./modules/codedeploy"
  project_name     = var.project_name
  service_role_arn = module.iam_roles.codedeploy_role_arn
  ec2_tag_key      = var.ec2_tag_key
  ec2_tag_value    = var.ec2_tag_value
}

module "codepipeline" {
  source                = "./modules/codepipeline"
  project_name          = var.project_name
  service_role_arn      = module.iam_roles.codepipeline_role_arn
  artifacts_bucket      = aws_s3_bucket.codepipeline_artifacts.bucket
  github_repo           = "${var.github_owner}/${var.github_repo}"
  github_branch         = var.github_branch
  github_connection_arn = var.github_connection_arn
  codebuild_project     = module.codebuild.project_name
  codedeploy_app        = module.codedeploy.application_name
  deployment_group      = module.codedeploy.deployment_group_name
} 