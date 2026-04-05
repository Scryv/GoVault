#!/bin/bash

set -euo pipefail
require_root() {
  if [[ $EUID -ne 0 ]]; then
    echo "Error: must run as root"
    exit 1
  fi
}

require_root
echo "Welcome to GoVault Installer"
go build
chmod +x GoVault
echo "made GoVault Executable"
sudo mv GoVault /usr/local/bin/
echo "Copied GoVault to /user/local/bin"
echo "Install is finished try it out by typing the command GoVault!"


