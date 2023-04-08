#!/bin/bash

if [ -d "$HOME/.config/nxm/bitwarden" ]; then
    continue
else
    mkdir -p $HOME/.config/nxm/bitwarden
fi

cd $HOME/.config/nxm/bitwarden

curl -Lso bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
./bitwarden.sh install

