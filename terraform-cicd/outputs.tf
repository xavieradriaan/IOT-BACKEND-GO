output "s3_bucket_name" {
  description = "Name of the S3 bucket for artifacts"
  value       = aws_s3_bucket.codepipeline_artifacts.bucket
}

output "pipeline_name" {
  description = "Name of the CodePipeline"
  value       = module.codepipeline.pipeline_name
}

output "codebuild_project_name" {
  description = "Name of the CodeBuild project"
  value       = module.codebuild.project_name
}

output "codedeploy_app_name" {
  description = "Name of the CodeDeploy application"
  value       = module.codedeploy.application_name
}

output "codedeploy_deployment_group" {
  description = "Name of the CodeDeploy deployment group"
  value       = module.codedeploy.deployment_group_name
} 