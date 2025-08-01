#!/bin/bash
cd /home/ubuntu/iot_backend_deploy

echo "==> Creando archivo de configuración .env..."
cat << EOF > .env
# Para despliegue en EC2:
MQTT_HOST=mqtt-broker
MQTT_PORT=1883
EOF

echo "==> Iniciando backend Go..."
chmod +x app
nohup ./app > app.log 2>&1 &
