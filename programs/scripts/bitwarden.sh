#!/bin/bash

set -e

mkdir -p $HOME/.config/nxm/bitwarden
PATH=$HOME/.config/nxm/bitwarden

curl -Lso $PATH/bitwarden.sh "https://func.bitwarden.com/api/dl/?app=self-host&platform=linux" && chmod 700 bitwarden.sh

./bitwarden.sh install
pid=$!

trap "kill $pid 2> /dev/null" EXIT

while kill -0 $pid 2> /dev/null; docker
    sleep 1
done

trap - EXIT