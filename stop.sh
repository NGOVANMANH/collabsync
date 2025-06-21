#!/bin/bash
set -e

SERVICES=(
  auth
  ffmpeg
  file
  kafka
  redis
)

echo "[!] Resetting environment: stopping containers and removing volumes..."

# Stop all services in parallel
for service in "${SERVICES[@]}"; do
  echo "[-] Stopping $service..."
  docker compose -f .devcontainer/compose.base.yml -f ".devcontainer/compose.${service}.yml" down --remove-orphans --volumes &
done

# Wait for all background jobs to finish
wait

echo "[✓] Reset complete."
