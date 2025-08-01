output "project_name" {
  value = aws_codebuild_project.main.name
}

output "project_arn" {
  value = aws_codebuild_project.main.arn
} 