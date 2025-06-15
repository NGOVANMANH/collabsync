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

# Reset mode: bring everything down before bootstrapping
echo "[!] Resetting environment: stopping containers and removing volumes..."
for service in "${SERVICES[@]}"; do
docker compose -f compose.base.yml -f "compose.${service}.yml" down || true
done
echo "[✓] Reset complete."

