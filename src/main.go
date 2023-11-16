package main

import (
    "fmt"
    "os"

    "yourproject/src/libb/keycard_link"
    "yourproject/src/libb/ipfs_link"
)

func main() {
	checkDependencies()
    // Example: Parsing command-line arguments
    // You might want to use a more robust way for parsing arguments (like the `flag` package)
    if len(os.Args) < 2 {
        fmt.Println("Usage: main.go [command]")
        os.Exit(1)
    }

    command := os.Args[1]

    switch command {
    case "encrypt":
        if len(os.Args) != 3 {
            fmt.Println("Usage: main.go encrypt [filename]")
            os.Exit(1)
        }
        filename := os.Args[2]
        err := encryptFile(filename)
        if err != nil {
            fmt.Printf("Error encrypting file: %s\n", err)
            os.Exit(1)
        }

    case "decrypt":
        if len(os.Args) != 3 {
            fmt.Println("Usage: main.go decrypt [filename]")
            os.Exit(1)
        }
        filename := os.Args[2]
        err := decryptFile(filename)
        if err != nil {
            fmt.Printf("Error decrypting file: %s\n", err)
            os.Exit(1)
        }

    default:
        fmt.Println("Invalid command")
        os.Exit(1)
    }
}
func askUserYN(question string) bool {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Printf("%s (y/n): ", question)
        response, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading response. Please try again.")
            continue
        }
        response = strings.TrimSpace(strings.ToLower(response))

        if response == "y" || response == "yes" {
            return true
        } else if response == "n" || response == "no" {
            return false
        } else {
            fmt.Println("Invalid response. Please answer 'y' for yes or 'n' for no.")
        }
    }
}

func checkAndInstallDependencies() error {
    dependencies := []string{"git", "ipfs"}

    for _, dep := range dependencies {
        _, err := exec.LookPath(dep)
        if err != nil {
            if askUserYN(fmt.Sprintf("'%s' is not installed. Do you want to install it?", dep)) {
                err := installDependency(dep)
                if err != nil {
                    return fmt.Errorf("error installing %s: %w", dep, err)
                }
            }
        }
    }
    return nil
}

func installDependency(dep string) error {
    // Implement the logic to install the dependency.
    // This is an example using apt-get, adjust according to your system's package manager.
    cmd := exec.Command("sudo", "pacman","-S", dep)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func encryptFile(filename string) error {
    encryptedFilePath, err := keycard_link.EncryptFile(filename)
    if err != nil {
        return fmt.Errorf("error encrypting file: %w", err)
    }
    fmt.Printf("File encrypted successfully: %s\n", encryptedFilePath)

    if askUserYN("Do you want to upload the file to IPFS? [y/n]: ") {
        cid, err := ipfs_link.AddFileToIPFS(encryptedFilePath)
        if err != nil {
            return fmt.Errorf("error uploading file to IPFS: %w", err)
        }
        fmt.Printf("File uploaded to IPFS with CID: %s\n", cid)
    }

    return nil // Return nil if successful, or an error if something goes wrong
}

func decryptFile(filename string) error {
    // Retrieve the file from IPFS using ipfs_link library
    err := ipfs_link.GetFileFromIPFS(/* ... */)
    if err != nil {
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }

    decryptedData, err := keycard_link.DecryptFile(/* ... */)
    if err != nil {
        return fmt.Errorf("error decrypting file: %w", err)
    }

    // Save the decrypted data to a file
    // ...

    return nil // Return nil if successful, or an error if something goes wrong
}




