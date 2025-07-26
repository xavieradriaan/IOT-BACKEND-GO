#!/bin/bash
echo "==> Deteniendo backend Go..."
pkill -f "./app" || true
