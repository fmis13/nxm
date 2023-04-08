#!/bin/bash

set -e

echo 'All important information will be stored at $HOME/home-assistant'

if [ ! -d "$HOME/.config/nxm/home-assistant" ] && [ ! -d "$HOME/.config/nxm/home-assistant/config" ]; then
    mkdir -p $HOME/.config/nxm/home-assistant
    mkdir $HOME/.config/nxm/home-assistant/config
fi

	cat > $HOME/.config/nxm/home-assistant/docker-compose.yml << END
version: '3'
services:
  homeassistant:
    container_name: homeassistant
    image: "ghcr.io/home-assistant/home-assistant:stable"
    volumes:
      - $HOME/.config/nxm/home-assistant/config:/config
      - /etc/localtime:/etc/localtime:ro
    restart: unless-stopped
    privileged: true
    network_mode: host
END

cd $HOME/.config/nxm/home-assistant

docker compose up -d