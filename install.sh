#!/bin/bash

# DPPM Installation Script
# Automatically detects platform and installs the appropriate binary

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# GitHub repository
REPO="AI-S-Tools/dp-project-app"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="dppm"

echo -e "${BLUE}🚀 DPPM (Dropbox Project Manager) Installer${NC}"
echo -e "${BLUE}=============================================${NC}"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
    linux*)
        OS="linux"
        ;;
    darwin*)
        OS="macos"
        ;;
    mingw* | msys* | cygwin*)
        OS="windows"
        ;;
    *)
        echo -e "${RED}❌ Unsupported operating system: $OS${NC}"
        exit 1
        ;;
esac

case $ARCH in
    x86_64 | amd64)
        ARCH="amd64"
        ;;
    aarch64 | arm64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}❌ Unsupported architecture: $ARCH${NC}"
        exit 1
        ;;
esac

# Construct binary name
if [ "$OS" = "windows" ]; then
    BINARY_FILE="dppm-${OS}-${ARCH}.exe"
else
    BINARY_FILE="dppm-${OS}-${ARCH}"
fi

echo -e "${YELLOW}📋 Detected platform: ${OS}-${ARCH}${NC}"
echo -e "${YELLOW}📦 Binary: ${BINARY_FILE}${NC}"

# Get latest release info
echo -e "${BLUE}🔍 Fetching latest release information...${NC}"
RELEASE_INFO=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest")
LATEST_VERSION=$(echo "$RELEASE_INFO" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo -e "${RED}❌ Failed to get latest version information${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Latest version: ${LATEST_VERSION}${NC}"

# Construct download URL
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_VERSION}/${BINARY_FILE}"

echo -e "${BLUE}⬇️  Downloading ${BINARY_FILE}...${NC}"

# Create temporary directory
TMP_DIR=$(mktemp -d)
TMP_FILE="${TMP_DIR}/${BINARY_FILE}"

# Download binary
if ! curl -L -o "$TMP_FILE" "$DOWNLOAD_URL"; then
    echo -e "${RED}❌ Failed to download binary${NC}"
    rm -rf "$TMP_DIR"
    exit 1
fi

echo -e "${GREEN}✅ Download completed${NC}"

# Make binary executable
chmod +x "$TMP_FILE"

# Install binary
echo -e "${BLUE}📁 Installing to ${INSTALL_DIR}/${BINARY_NAME}...${NC}"

if [ "$OS" = "windows" ]; then
    # Windows installation
    if ! cp "$TMP_FILE" "${INSTALL_DIR}/${BINARY_NAME}.exe" 2>/dev/null; then
        echo -e "${YELLOW}⚠️  Failed to install to ${INSTALL_DIR}. Installing to user directory...${NC}"
        USER_BIN_DIR="$HOME/bin"
        mkdir -p "$USER_BIN_DIR"
        cp "$TMP_FILE" "${USER_BIN_DIR}/${BINARY_NAME}.exe"
        echo -e "${GREEN}✅ Installed ${BINARY_NAME}.exe to ${USER_BIN_DIR}${NC}"

        # Add to PATH in Windows
        echo -e "${BLUE}🔧 Adding to PATH...${NC}"
        echo "set PATH=%PATH%;%USERPROFILE%\\bin" >> ~/.bashrc 2>/dev/null || true
        echo -e "${YELLOW}📝 Added to PATH. Restart your terminal or run: set PATH=%PATH%;%USERPROFILE%\\bin${NC}"
    else
        echo -e "${GREEN}✅ Successfully installed to ${INSTALL_DIR}/${BINARY_NAME}.exe${NC}"
    fi
else
    # Unix-like systems (Linux, macOS)
    if sudo cp "$TMP_FILE" "${INSTALL_DIR}/${BINARY_NAME}" 2>/dev/null; then
        echo -e "${GREEN}✅ Successfully installed to ${INSTALL_DIR}/${BINARY_NAME}${NC}"
    else
        echo -e "${YELLOW}⚠️  No sudo access. Installing to user directory...${NC}"
        USER_BIN_DIR="$HOME/.local/bin"
        mkdir -p "$USER_BIN_DIR"
        cp "$TMP_FILE" "${USER_BIN_DIR}/${BINARY_NAME}"
        echo -e "${GREEN}✅ Installed to ${USER_BIN_DIR}/${BINARY_NAME}${NC}"

        # Add to PATH automatically
        echo -e "${BLUE}🔧 Adding to PATH...${NC}"

        # Detect shell and add to appropriate config file
        SHELL_NAME=$(basename "$SHELL")
        case $SHELL_NAME in
            bash)
                SHELL_CONFIG="$HOME/.bashrc"
                ;;
            zsh)
                SHELL_CONFIG="$HOME/.zshrc"
                ;;
            fish)
                SHELL_CONFIG="$HOME/.config/fish/config.fish"
                ;;
            *)
                SHELL_CONFIG="$HOME/.profile"
                ;;
        esac

        # Check if PATH is already in config
        PATH_LINE="export PATH=\"\$HOME/.local/bin:\$PATH\""
        if [ "$SHELL_NAME" = "fish" ]; then
            PATH_LINE="set -gx PATH \$HOME/.local/bin \$PATH"
        fi

        if ! grep -q "$HOME/.local/bin" "$SHELL_CONFIG" 2>/dev/null; then
            echo "" >> "$SHELL_CONFIG"
            echo "# Added by DPPM installer" >> "$SHELL_CONFIG"
            echo "$PATH_LINE" >> "$SHELL_CONFIG"
            echo -e "${GREEN}✅ Added to PATH in $SHELL_CONFIG${NC}"
            echo -e "${YELLOW}📝 Restart your terminal or run: source $SHELL_CONFIG${NC}"
        else
            echo -e "${BLUE}ℹ️  PATH already configured in $SHELL_CONFIG${NC}"
        fi

        # Also add to current session
        export PATH="$HOME/.local/bin:$PATH"
        echo -e "${GREEN}✅ Added to current session PATH${NC}"
    fi
fi

# Cleanup
rm -rf "$TMP_DIR"

echo -e "${GREEN}🎉 Installation completed successfully!${NC}"

# Test installation
echo -e "${BLUE}🧪 Testing installation...${NC}"
if command -v $BINARY_NAME >/dev/null 2>&1; then
    VERSION_OUTPUT=$(${BINARY_NAME} --version 2>/dev/null || echo "")
    if [ -n "$VERSION_OUTPUT" ]; then
        echo -e "${GREEN}✅ Installation verified successfully!${NC}"
        echo -e "${GREEN}   $VERSION_OUTPUT${NC}" | head -1
    else
        echo -e "${YELLOW}⚠️  Binary installed but unable to get version${NC}"
    fi
else
    echo -e "${YELLOW}⚠️  Binary installed but not in current PATH${NC}"
    echo -e "${YELLOW}   Please restart your terminal or run: source ~/.bashrc${NC}"
fi

echo
echo -e "${BLUE}🚀 Getting Started:${NC}"
echo -e "  ${BINARY_NAME} --version                # Check installation"
echo -e "  ${BINARY_NAME}                         # Smart startup guide"
echo -e "  ${BINARY_NAME} wiki \"getting started\"   # Quick start guide"
echo -e "  ${BINARY_NAME} wiki list               # See all available topics"
echo
echo -e "${BLUE}📖 Learn More:${NC}"
echo -e "  Repository: https://github.com/${REPO}"
echo -e "  Documentation: Use the built-in wiki system"
echo
echo -e "${GREEN}Happy project managing with DPPM! 🎯${NC}"

# Quick start hint
echo
echo -e "${BLUE}💡 Quick Start Hint:${NC}"
echo -e "   Try: ${BINARY_NAME} wiki \"ai workflow\" to see AI-optimized usage patterns"