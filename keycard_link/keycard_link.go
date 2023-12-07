package keycard_link

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "io"
    "runtime"
    "os/exec"
    "strings"
    "net/http"
    "path/filepath"
)

// GLOBAL PLATFORM PRO VERSION V20.01.23

const gpJarURL = "https://github.com/martinpaljak/GlobalPlatformPro/releases/download/v20.01.23/gp.jar"

func InstallKeycard() error {
	cmd := exec.Command("./keycard-linux-amd64", "install")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing keycard-linux-amd64 install: %v, output: %s", err, output)
	}
	fmt.Printf("Output: %s\n", output)
	return nil
}

func GetKeycardPublicKey() (string, error) {
    fmt.Println("\033[1;33m==============================\033[0m")
    fmt.Println("\033[1;32m>\033[0m Scan your card now!")
    fmt.Println("\033[1;33m==============================\033[0m")
    cmd := exec.Command("./keycard-linux-amd64", "info")

    var out bytes.Buffer
    cmd.Stdout = &out

    // Run the command
    err := cmd.Run()
    if err != nil {
        return "", err
    }

    scanner := bufio.NewScanner(&out)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "PublicKey:") {
            parts := strings.Fields(line)
            if len(parts) > 1 {
                return parts[len(parts)-1], nil
            }
        }
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

    return "", fmt.Errorf("public key not found in the output")
}

// ReadPassphrase prompts the user to enter a passphrase.
func ReadPassphrase() (string, error) {

    fmt.Println("\033[1;33m==============================\033[0m")
    fmt.Println("\033[1;32m>\033[0m Enter a unique passphrase for this particular file: \n")
    fmt.Println("\033[1;33m==============================\033[0m")


    reader := bufio.NewReader(os.Stdin)
    passphrase, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(passphrase), nil
}




func JavaDependency() {
	// Check if Java is installed
	if err := exec.Command("java", "-version").Run(); err == nil {
		fmt.Println("Java is already installed! :)")
		return
	}

	// Determine the operating system
	switch os := runtime.GOOS; os {
	case "linux":
		// Identify Linux distribution
		out, err := exec.Command("sh", "-c", "cat /etc/os-release").Output()
		if err != nil {
			fmt.Printf("Failed to execute command: %s\n", err)
			return
		}

		// Parse /etc/os-release
		osRelease := string(out)
		if strings.Contains(osRelease, "ID=arch") || strings.Contains(osRelease, "ID_LIKE=arch") {
			// Install Java for Arch-based systems
			fmt.Println("Installing Java using pacman...")
			exec.Command("sudo", "pacman", "-Sy", "--noconfirm", "jdk-openjdk").Run()

			// Install pcscd and pcsc-tools using pacman
			fmt.Println("Installing pcscd and pcsc-tools using pacman...")
			exec.Command("sudo", "pacman", "-Syu", "--noconfirm", "ccid", "libnfc", "acsccid", "pcsclite", "pcsc-tools").Run()
			exec.Command("sudo", "systemctl", "enable", "--now", "pcscd").Run()
		} else if strings.Contains(osRelease, "ID=debian") || strings.Contains(osRelease, "ID=ubuntu") {
			// Install Java for Debian-based systems
			fmt.Println("Installing Java using apt...")
			exec.Command("sudo", "apt", "update").Run()
			exec.Command("sudo", "apt", "install", "-y", "default-jdk").Run()

			// Install pcscd and pcsc-tools using apt
			fmt.Println("Installing pcscd and pcsc-tools using apt...")
			exec.Command("sudo", "apt-get", "update").Run()
			exec.Command("sudo", "apt-get", "install", "-y", "pcsc-lite", "pcsc-tools").Run()
			exec.Command("sudo", "systemctl", "enable", "pcscd").Run()
			exec.Command("sudo", "systemctl", "start", "pcscd").Run()
		} else {
			fmt.Println("Unsupported Linux distribution")
		}
	default:
		fmt.Printf("%s is not supported by this script\n", os)
	}
}

func GlobalPlatformDependency() {
	// Define paths to search
	paths := []string{"../", os.Getenv("HOME") + "/Downloads/"}

	found := false
	for _, path := range paths {
		found = searchFile(path, "gp.jar")
		if found {
			fmt.Println("gp.jar found in:", path)
			fmt.Println("Nice, GlobalPlatformDependency is good :) ")
			break
		}
	}

	// If not found, download the file
	if !found {
		fmt.Println("gp.jar not found :/  Downloading...")
		err := downloadFile("gp.jar", gpJarURL)
		if err != nil {
			fmt.Println("Error downloading gp.jar:", err)
			return
		}
		fmt.Println("Downloaded gp.jar successfully.")
	}
}

func searchFile(dir, filename string) bool {
	found := false
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == filename {
			found = true
		}
		return nil
	})
	return found
}





func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	buf := make([]byte, 1024) // buffer size
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, writeErr := out.Write(buf[:n]); writeErr != nil {
			return writeErr
		}
	}

	return nil
}


