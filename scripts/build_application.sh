#!/bin/bash
set -e

echo "==> DEBUG: Mostrando directorio actual y contenido"
pwd
ls -la

echo "==> DEBUG: Buscando archivos app en todo el sistema"
find /opt/codedeploy-agent/deployment-root -name "app" -type f 2>/dev/null || echo "No se encontró archivo app"

echo "==> DEBUG: Listando contenido de deployment-archive"
ls -la /opt/codedeploy-agent/deployment-root/*/*/deployment-archive/ 2>/dev/null || echo "No se pudo acceder a deployment-archive"

cd /home/ubuntu/iot_backend_deploy

echo "==> Tomando propiedad del directorio de trabajo"
sudo chown -R ubuntu:ubuntu .

echo "==> Verificando archivos de aplicación"
ls -la

echo "==> DEBUG: Verificando si app existe antes de chmod"
if [ -f app ]; then
    echo "==> app encontrado, aplicando chmod"
    chmod +x app
else
    echo "==> ERROR: app no encontrado en directorio de trabajo"
    echo "==> Listando todo el contenido:"
    find . -type f
fi

echo "==> Configuración de aplicación completada exitosamente"