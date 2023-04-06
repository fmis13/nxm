#!/bin/bash

adduser bitwarden --disabled-password
usermod -aG docker bitwarden
mkdir /opt/bitwarden
chmod -R 700 /opt/bitwarden
chown -R bitwarden:bitwarden /opt/bitwarden
runuser -l bitwarden -c 'curl -Lso bitwarden.sh https://go.btwrdn.co/bw-sh && chmod 700 bitwarden.sh'
runuser -l bitwarden -c './bitwarden.sh install'
runuser -l bitwarden -c './bitwarden.sh start'