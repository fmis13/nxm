#!/bin/bash

set -e

mkdir -p $HOME/.config/nxm/bitwarden
cd $HOME/.config/nxm/bitwarden

curl -Lso bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
while true; do
{
    ./bitwarden.sh install
    sleep 60
}