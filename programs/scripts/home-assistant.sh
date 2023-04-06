#!/bin/bash

set -e

echo 'All important information will be stored at $HOME/home-assistant'

mkdir $HOME/home-assistant
mkdir $HOME/home-assistant/config

	cat > $HOME/home-assistant/docker-compose.yml << END
version: '3'
services:
  homeassistant:
    container_name: homeassistant
    image: "ghcr.io/home-assistant/home-assistant:stable"
    volumes:
      - /PATH_TO_YOUR_CONFIG:/config
      - /etc/localtime:/etc/localtime:ro
    restart: unless-stopped
    privileged: true
    network_mode: host
END

cd $HOME/home-assistant
docker compose up -d