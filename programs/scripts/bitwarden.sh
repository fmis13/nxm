#!/bin/bash

set -e

mkdir -p $HOME/.config/nxm/bitwarden
PATH=$HOME/.config/nxm/bitwarden

curl -Lso $PATH/bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
echo "Run $HOME/.config/nxm/bitwarden/bitwarden.sh install to finish the installation"