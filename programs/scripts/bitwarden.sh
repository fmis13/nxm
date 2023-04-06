#!/bin/bash

echo "Hi"


mkdir -p $HOME/.config/nxm/bitwarden
cd $HOME/.config/nxm/bitwarden
echo "Hi"

curl -Lso bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
echo "Hi"
./bitwarden.sh install

