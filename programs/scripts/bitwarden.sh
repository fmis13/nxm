#!/bin/bash

set -e
set +x

mkdir -p $HOME/.config/nxm/bitwarden
BWPATH=$HOME/.config/nxm/bitwarden

curl -Lso $BWPATH/bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh
$BWPATH/bitwarden.sh install