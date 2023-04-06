#!/bin/bash

set -e
set +x

mkdir -p $HOME/.config/nxm/bitwarden
PATH=$PATH:$HOME/.config/nxm/bitwarden

curl -Lso $PATH/bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
$PATH/bitwarden.sh install