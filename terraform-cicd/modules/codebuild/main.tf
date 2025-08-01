resource "aws_codebuild_project" "main" {
  name          = "${var.project_name}-build"
  description   = "Build project for ${var.project_name}"
  build_timeout = "10"
  service_role  = var.service_role_arn

  artifacts {
    type = "CODEPIPELINE"
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL" 
    image                       = "aws/codebuild/amazonlinux2-x86_64-standard:5.0"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"
    privileged_mode             = true

    environment_variable {
      name  = "GOOS"
      value = "linux"
    }

    environment_variable {
      name  = "GOARCH"
      value = "amd64"
    }

    environment_variable {
      name  = "CGO_ENABLED"
      value = "0"
    }
  }

  source {
    type = "CODEPIPELINE"
    buildspec = yamlencode({
      version = "0.2"
      phases = {
        pre_build = {
          commands = [
            "echo Logging in to Amazon ECR...",
            "aws --version",
            "go version"
          ]
        }
        build = {
          commands = [
            "echo Build started on `date`",
            "go mod tidy",
            "go build -o app main.go",
            "echo Build completed on `date`"
          ]
        }
        post_build = {
          commands = [
            "echo Build completed on `date`",
            "ls -la"
          ]
        }
      }
      artifacts = {
        files = [
          "app",
          "appspec.yml",
          "scripts/**/*",
          "init_users.sql"
        ]
      }
    })
  }

  logs_config {
    cloudwatch_logs {
      group_name  = "/aws/codebuild/${var.project_name}"
      stream_name = "build-log"
    }
  }
} 