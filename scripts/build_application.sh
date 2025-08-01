#!/bin/bash
set -e

cd /home/ubuntu/iot_backend_deploy

echo "==> Tomando propiedad del directorio de trabajo"
sudo chown -R ubuntu:ubuntu .

echo "==> Verificando archivos de aplicación"
ls -la

echo "==> El binario ya fue compilado en CodeBuild"
chmod +x app

echo "==> Configuración de aplicación completada exitosamente"