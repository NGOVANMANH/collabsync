#!/bin/bash
set -e

SERVICES=(
  auth
  ffmpeg
  file
  kafka
  redis
)

echo "==========================================="
echo "🔧 Bootstrapping Docker Environment: collabsync"
echo "==========================================="

# Merge .env files
./copyenv.sh

# Create network only if not exists
NETWORK_NAME="collabsync-network"
echo "[+] Creating network..."
if ! docker network ls --format '{{.Name}}' | grep -q "^${NETWORK_NAME}$"; then
  echo "[+] Creating Docker network: ${NETWORK_NAME}"
  docker network create "${NETWORK_NAME}"
else
  echo "[✓] Docker network already exists: ${NETWORK_NAME}"
fi

# Start all services in parallel
for service in "${SERVICES[@]}"; do
  echo "[+] Starting $service service..."
  docker compose \
    -f .devcontainer/compose.base.yml \
    -f .devcontainer/compose.${service}.yml \
    up -d --remove-orphans &
done

# Wait for all background jobs
wait

echo "==========================================="
echo "🚀 Docker environment bootstrapped successfully!"
echo "💡 You can now access all running services."
echo "==========================================="
