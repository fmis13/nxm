#!/bin/bash
# curl https://github.com/nixmember/nixmember/raw/master/install.sh | sh

# A bash script that installs the latest release
# Mostly taken from https://raw.githubusercontent.com/srevinsaju/zap/main/install.sh - thank you Srevin, you have been a big inspiration

# pipefail
set -euo pipefail

if [ -z "$(which figlet)" ]; then
	echo "nxm"
	echo "The package manager for linux (home-)servers."
else
	figlet "nxm"
	echo "The package manager for linux (home-)servers."
fi

# env
REPO="nixmember/nixmember"

if [ -z ${ARCH+x} ]; then
	MACHINE_ARCH="$(uname -m)"
	if [ "$MACHINE_ARCH" = "amd64" ]; then
		ARCH="amd64"
        elif [ "$MACHINE_ARCH" = "x86_64" ]; then
                ARCH="amd64"
	elif [ "$MACHINE_ARCH" = "aarch64" ]; then
		ARCH="arm64"
	elif [ "$MACHINE_ARCH" = "arm64" ]; then
		ARCH="arm64"
	fi
    export ARCH
fi

echo
echo "—"
echo

echo "Checking if all required programs are installed: "

# required: curl
if ! command -v curl &>/dev/null; then
	echo
	echo curl is required and was not found. Please install curl.
	exit 1
else
	echo -e [OK] curl
fi

# required: grep
if ! command -v grep &>/dev/null; then
	echo
	echo grep is required and was not found. Please install grep.
	exit 1
else
	echo -e [OK] grep
fi


# required: jq

if ! command -v jq &>/dev/null; then
	echo
	echo jq is required and was not found. Please install jq.
	echo
	exit 1
else
	echo -e [OK] jq
fi

RELEASE_API_URL="https://api.github.com/repos/$REPO/releases"

echo
echo Grabbing latest release....

RELEASES="$(curl -sN $RELEASE_API_URL)"

# List releases

COMPATIBLE_RELEASE="$(echo "$RELEASES" | jq -rc '.[].assets[].browser_download_url' | grep -m 1 "$ARCH")"

if [ -z "$COMPATIBLE_RELEASE" ]; then
	echo [!] No compatible releases found....
	exit 1
fi


echo

echo Found latest release...
echo

TMP="$(mktemp /tmp/nxm.XXXXXXXXXX)"

# Download release
echo Downloading....
echo "URL: $COMPATIBLE_RELEASE"
echo

curl -L "$COMPATIBLE_RELEASE" --output $TMP
echo

if [ "$EUID" -eq 0 ]; then
	echo Starting system-wide installation
	echo
	sudo mv "$TMP" /usr/bin/nxm
	sudo chmod +x /usr/bin/nxm
	echo "nxm installed to /usr/bin/nxm"
else
	echo Starting local installation
	echo
	echo If the installation errors out, make sure that ~/.local/bin exists.
	echo
	mv "$TMP" ~/.local/bin/nxm
	chmod +x ~/.local/bin/nxm
	echo "nxm installed to ~/.local/bin/nxm"
fi

echo
echo "Installation complete"