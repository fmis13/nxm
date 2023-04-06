#!/bin/bash

mkdir $HOME/.config/nxm/bitwarden
cd $HOME/.config/nxm/bitwarden

curl -Lso bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
cd $HOME/.config/nxm/bitwarden
echo "Please run ./bitwarden.sh install to finish the installation."
exit