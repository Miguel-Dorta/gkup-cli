#!/bin/bash
TMP_DIR="/tmp/Miguel-Dorta_gkup_install"
INSTALL_DIR="/opt/Miguel-Dorta/gkup"
ARCH="$(uname -m)"

case $ARCH in
x86_64)
  ARCH="amd64"
  ;;
i?86)
  ARCH="386"
  ;;
esac

mkdir $TMP_DIR || exit 1
cd $TMP_DIR || exit 1
wget -o gkup-core "https://github.com/Miguel-Dorta/gkup-core/releases/latest/download/gkup-core_$(uname -s)_$ARCH" || exit 1
wget -o gkup-cli "https://github.com/Miguel-Dorta/gkup-cli/releases/latest/download/gkup-cli_$(uname -s)_$ARCH" || exit 1

sudo mkdir $INSTALL_DIR
sudo mv gkup-core gkup-cli $INSTALL_DIR
sudo chown -R root:root $INSTALL_DIR
sudo chmod -R 0755 $INSTALL_DIR

rm -Rf $TMP_DIR
