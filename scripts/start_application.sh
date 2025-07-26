#!/bin/bash
cd /home/ubuntu/iot_backend_deploy
echo "==> Iniciando backend Go..."
chmod +x app
nohup ./app > app.log 2>&1 &
