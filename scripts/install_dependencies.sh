#!/bin/bash
set -e

cd /home/ubuntu/iot_backend_deploy

echo "==> Tomando propiedad del directorio de trabajo"
sudo chown -R ubuntu:ubuntu .

echo "==> Verificando instalación de Go"
which go || sudo apt-get update -y && sudo apt-get install -y golang

echo "==> Compilando binario Go"
go mod tidy
go build -o app main.go

echo "==> Instalación completada exitosamente"
