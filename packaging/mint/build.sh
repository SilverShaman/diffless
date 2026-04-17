#!/usr/bin/env bash
set -e

VERSION="1.0.0"
ARCH=$(dpkg --print-architecture 2>/dev/null || echo "amd64")
PKG_NAME="diffless_${VERSION}_${ARCH}"

echo "=> Building binary for Linux Mint (Debian pkg)..."
export PATH=$PATH:/usr/local/go/bin

# Compile Go from project root
cd ../../
go build -o packaging/mint/diffless_bin ./cmd/diffless/main.go

echo "=> Setting up Debian package structure..."
cd packaging/mint
mkdir -p "${PKG_NAME}/DEBIAN"
mkdir -p "${PKG_NAME}/usr/local/bin"

# Move binary
mv diffless_bin "${PKG_NAME}/usr/local/bin/diffless"
chmod +x "${PKG_NAME}/usr/local/bin/diffless"

# Create control file
cat <<EOF > "${PKG_NAME}/DEBIAN/control"
Package: diffless
Version: ${VERSION}
Section: standard
Priority: optional
Architecture: ${ARCH}
Maintainer: SilverShaman
Description: Diffless workflow CLI
 Designed to give AI agents physical sandbox isolation using git worktrees.
EOF

echo "=> Building .deb package..."
dpkg-deb --build "${PKG_NAME}"

echo "=> Success! Built ${PKG_NAME}.deb"
