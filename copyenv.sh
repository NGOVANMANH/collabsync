#!/bin/bash

TARGET_ENV_FILE="./.devcontainer/.env"
SOURCE_DIR="./src"

# Enable recursive globbing
shopt -s globstar nullglob

# Clear or create the target .env file
echo "# Auto-generated merged .env file" > "$TARGET_ENV_FILE"
echo "" >> "$TARGET_ENV_FILE"

# Loop through all .env files under ./src/**/
for env_file in "$SOURCE_DIR"/**/.env; do
    if [ -f "$env_file" ]; then
        service_name=$(basename "$(dirname "$env_file")")
        echo "[+] Merging $env_file into $TARGET_ENV_FILE"

        {
            echo "# ----- $service_name -----"
            cat "$env_file"
            echo ""
        } >> "$TARGET_ENV_FILE"
    else
        echo "[✖] No .env file found in $env_file"
    fi
done

echo "[✔] Merged all .env files into $TARGET_ENV_FILE"
