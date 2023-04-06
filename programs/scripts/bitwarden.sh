#!/bin/bash

set -e

mkdir -p $HOME/.config/nxm/bitwarden
cd $HOME/.config/nxm/bitwarden

curl -Lso bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
./bitwarden.sh install
while true; do
    if ! pgrep bash; then
        break
    fi
done