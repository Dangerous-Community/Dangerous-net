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
            sudo apt-get install pcscd pcsc-tools
            ;;
        2)
            echo "Installing pcscd and pcsc-tools using pacman..."
            sudo pacman -Sy pcscd pcsc-tools
            ;;
        3)
            read -p "Write install command for pcsc-tools or download similar package: " custom_install_cmd
            eval $custom_install_cmd
            ;;
        *)
            echo "Invalid choice. Exiting."
            exit 1
            ;;
    esac
}

blacklist_drivers() {
    echo "Blacklisting conflicting NFC kernel drivers..."
    echo -e "install nfc /bin/false\ninstall pn533 /bin/false" | sudo tee -a /etc/modprobe.d/blacklist.conf
}

download_install_drivers() {
    echo "Downloading ACR122U drivers..."
    wget http://www.acs.com.hk/download-driver-unified/11929/ACS-Unified-PKG-Lnx-118-P.zip -O acr122u_driver.zip

    echo "Unzipping downloaded drivers..."
    unzip acr122u_driver.zip -d acr122u_driver

    echo "Installing drivers..."
    cd acr122u_driver
    # Placeholder for driver installation command
    # ./install.sh or other commands as per the driver documentation
}

# Run functions
install_pcsc
blacklist_drivers
download_install_drivers

echo "Please follow any on-screen instructions to complete the driver installation."

echo "Once the driver is installed, restart your computer."

echo "After restarting, run 'pcsc_scan' to verify the installation. You should see 'ACS ACR122U' listed among the devices."

