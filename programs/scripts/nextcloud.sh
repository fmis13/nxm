#!/bin/bash

set -e

read -p 'Choose a folder to store all your data: ' FOLDER

do 
    if [ -d "$HOME/.config/nxm/nextcloud" ]; then
        continue
    fi
    mkdir -p $HOME/.config/nxm/nextcloud
done

docker volume create \
--driver local \
--name nextcloud_aio_nextcloud_datadir \
-o device=$FOLDER \
-o type="none" \
-o o="bind"

docker run \
--sig-proxy=false \
--name nextcloud-aio-mastercontainer \
--restart always \
--publish 80:80 \
--publish 8080:8080 \
--publish 8443:8443 \
--volume nextcloud_aio_mastercontainer:/mnt/docker-aio-config \
--volume /var/run/docker.sock:/var/run/docker.sock:ro \
-e NEXTCLOUD_DATADIR="nextcloud_aio_nextcloud_datadir"
nextcloud/all-in-one:latest