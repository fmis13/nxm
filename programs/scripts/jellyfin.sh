#!/bin/bash

set -e

read -p 'Choose the folder where your media is stored (please use the standard linux paths, e.g /mnt/myvideos): ' FOLDER
echo "You selected folder $FOLDER"

	echo 'Sve informacije o ovome programu bit će spremljene u $HOME/jellyfin'
	mkdir $HOME/jellyfin
	mkdir $HOME/jellyfin/config
	mkdir $HOME/jellyfin/cache

  mkdir -p $HOME/.config/nxm/jellyfin
  mkdir $HOME/.config/nxm/jellyfin/config
  mkdir $HOME/.config/nxm/jellyfin/cache


	cat > $HOME/.config/nxm/jellyfin/docker-compose.yml << END
version: '3.5'
services:
  jellyfin:
    image: jellyfin/jellyfin
    container_name: jellyfin
    user: uid:gid
    network_mode: 'host'
    volumes:
      - /path/to/config:/config
      - /path/to/cache:/cache
      - "$FOLDER":/media
    restart: 'unless-stopped'
    # Optional - alternative address used for autodiscovery
    environment:
      - JELLYFIN_PublishedServerUrl=http://example.com
    # Optional - may be necessary for docker healthcheck to pass if running in host network mode
    extra_hosts:
      - "host.docker.internal:host-gateway"
END
	cd $HOME/.config/nxm/jellyfin
	docker compose up -d