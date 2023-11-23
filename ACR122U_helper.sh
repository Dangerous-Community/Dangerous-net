#!/bin/bash

echo "Select your Linux distribution:"
echo "1) Debian-based (Ubuntu, Linux Mint, etc.)"
echo "2) Arch-based (Arch Linux, Manjaro, etc.)"
echo "3) Other"
read -p "Enter your choice (1, 2, or 3): " distro_choice

install_pcsc() {
    case $distro_choice in
        1)
            echo "Installing pcscd and pcsc-tools using apt..."
            sudo apt-get update
            sudo apt-get install -y pcsc-lite pcsc-tools
            sudo systemctl enable pcscd
            sudo systemctl start pcscd
            ;;
        2)
            echo "Installing pcscd and pcsc-tools using pacman..."
            sudo pacman -Syu --noconfirm ccid libnfc acsccid pcsclite pcsc-tools
            sudo systemctl enable --now pcscd
            ;;
        3)
            read -p "Enter the install command for pcsc-tools or download a similar package: " custom_install_cmd
            eval $custom_install_cmd
            ;;
        *)
            echo "Invalid choice. Exiting."
            exit 1
            ;;
    esac
}

echo "Please ensure that your smart card reader is properly connected."
echo "Running pcsc_scan to detect the reader:"

# Check if pcsc_scan is available
if ! command -v pcsc_scan &>/dev/null; then
    echo "pcsc_scan command not found. Please install pcsc-tools."
    exit 1
fi

pcsc_scan

# Run functions
install_pcsc
