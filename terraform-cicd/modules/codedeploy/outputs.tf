output "application_name" {
  value = aws_codedeploy_app.main.name
}

output "deployment_group_name" {
  value = aws_codedeploy_deployment_group.main.deployment_group_name
} 