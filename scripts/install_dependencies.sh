#!/bin/bash
set -e

echo "==> Creando directorio de trabajo"
sudo mkdir -p /home/ubuntu/iot_backend_deploy
sudo chown -R ubuntu:ubuntu /home/ubuntu/iot_backend_deploy

echo "==> Verificando instalación de Go"
which go || (sudo apt-get update -y && sudo apt-get install -y golang-go)

echo "==> Instalación de dependencias completada exitosamente"
