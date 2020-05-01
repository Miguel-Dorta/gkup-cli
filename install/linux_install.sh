#!/bin/bash
TMP_DIR="/tmp/Miguel-Dorta_gkup_install"
INSTALL_DIR="/opt/Miguel-Dorta/gkup"
ARCH="$(uname -m)"
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"

case $ARCH in
x86_64)
  ARCH="amd64"
  ;;
i?86)
  ARCH="386"
  ;;
esac

# Download binaries
mkdir $TMP_DIR || exit 1
wget -qO- "https://github.com/Miguel-Dorta/gkup-cli/releases/latest/download/gkup_${OS}_${ARCH}.gz" | gzip -d > $TMP_DIR/gkup-cli || exit 1
wget -qO- "https://github.com/Miguel-Dorta/gkup-core/releases/latest/download/gkup-core_${OS}_${ARCH}.gz" | gzip -d > $TMP_DIR/gkup-core || exit 1

# Put binaries
sudo mkdir -p $INSTALL_DIR || exit 1
sudo mv $TMP_DIR/gkup-core $TMP_DIR/gkup-cli $INSTALL_DIR || exit 1
sudo ln -s $INSTALL_DIR/gkup-cli $INSTALL_DIR/gkup
sudo chown -R root:root $INSTALL_DIR || exit 1
sudo chmod -R 0755 $INSTALL_DIR || exit 1

# Clean up tmp dir
rm -Rf $TMP_DIR

# Set env vars and path
gkupPath="$INSTALL_DIR/gkup-core"
sudo bash -c "echo -e 'export GKUP_PATH=$INSTALL_DIR/gkup-core\nexport PATH=\$PATH:$INSTALL_DIR' > /etc/profile.d/gkup.sh"
sudo chmod 0755 $INSTALL_DIR/gkup-core
