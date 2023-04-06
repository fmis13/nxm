#!/bin/bash

set -e

mkdir -p $HOME/.config/nxm/bitwarden
PATH=$HOME/.config/nxm/bitwarden

curl -Lso $PATH/bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
echo "Please run ./bitwarden.sh install from the command line and fill in all the prompts, then follow the guide at https://bitwarden.com/help/install-on-premise-linux/"
$SHELL