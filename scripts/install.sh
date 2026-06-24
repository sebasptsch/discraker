#!/usr/bin/env bash
set -euo pipefail

REPO="sebasptsch/discraker"
APP="discraker"
INSTALL_DIR="$HOME/printer_data/tools/discraker"
SERVICE_NAME="discraker"

API_URL="https://api.github.com/repos/${REPO}/releases/latest"

# Determine GoReleaser asset name
OS="$(uname -s)"

case "$(uname -m)" in
    x86_64) ARCH="x86_64" ;;
    i386|i686) ARCH="i386" ;;
    aarch64) ARCH="arm64" ;;
    armv7l) ARCH="armv7" ;;
    armv6l) ARCH="armv6" ;;
    *) ARCH="$(uname -m)" ;;
esac

ASSET_NAME="discraker_${OS}_${ARCH}"

if [ "$EUID" -eq 0 ]; then
    SUDO=""
else
    SUDO="sudo"
fi

echo "Finding latest release..."

DOWNLOAD_URL="$(wget -qO- "$API_URL" \
    | grep browser_download_url \
    | grep "$ASSET_NAME\"" \
    | cut -d '"' -f 4)"

VERSION="$(wget -qO- "$API_URL" \
    | grep '"tag_name"' \
    | sed -E 's/.*"([^"]+)".*/\1/')"

if [ -z "$DOWNLOAD_URL" ]; then
    echo "Failed to find asset $ASSET_NAME in latest release"
    exit 1
fi

echo "Downloading $DOWNLOAD_URL"

install -d -m 755 "$INSTALL_DIR"

wget -O "${INSTALL_DIR}/${ASSET_NAME}" "$DOWNLOAD_URL"
$SUDO chmod +x "${INSTALL_DIR}/${ASSET_NAME}"

cat > "${INSTALL_DIR}/release_info.json" <<EOF
{
  "project_name": "discraker",
  "project_owner": "sebasptsch",
  "version": "${VERSION}",
  "asset_name": "${ASSET_NAME}"
}
EOF

$SUDO tee "/etc/systemd/system/${SERVICE_NAME}.service" >/dev/null <<EOF
[Unit]
Description=Discraker
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
WorkingDirectory=${INSTALL_DIR}
ExecStart=${INSTALL_DIR}/${ASSET_NAME}
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

$SUDO systemctl daemon-reload
$SUDO systemctl enable --now "${SERVICE_NAME}"

echo "Installed ${APP} ${VERSION}"
echo "Moonraker update manager config:"
echo
cat <<EOF
[update_manager discraker]
type: executable
channel: stable
repo: sebasptsch/discraker
path: ${INSTALL_DIR}
is_system_service: True
managed_services: discraker
EOF