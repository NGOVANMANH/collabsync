#!/bin/bash
set -e

# List of docker compose files for services
SERVICES=(
  auth
#   email
  ffmpeg
  file
  kafka
  redis
)

echo "==========================================="
echo "🔧 Bootstrapping Docker Environment: collabsync"
echo "==========================================="

# Start networks
echo "[+] Creating network..."
# Create network if it doesn't exist
if ! docker network ls --format '{{.Name}}' | grep -q "^collabsync-network$"; then
  echo "[+] Creating Docker network: collabsync-network"
  docker network create collabsync-network
else
  echo "[✓] Docker network already exists: collabsync-network"
fi

# Start services
for service in "${SERVICES[@]}"; do
  echo "[+] Starting $service service..."
  docker compose -f compose.base.yml -f "compose.${service}.yml" up -d
done

echo "==========================================="
echo "🚀 Docker environment bootstrapped successfully!"
echo "💡 You can now access all running services."
echo "==========================================="
