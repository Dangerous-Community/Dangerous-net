package main

import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "bufio"
    "net/http"
    "crypto/sha256"
    "strings"
    "IPFSS_IPFS-Secure/keycard_link"
    "IPFSS_IPFS-Secure/ipfs_link"
)

func main() {
    if !checkKeycardBinaryExists() {
        fmt.Println("Keycard binary not found. Downloading...")
        err := downloadKeycardBinary()
        if err != nil {
            fmt.Printf("Failed to download keycard binary: %s\n", err)
            os.Exit(1)
        }
        fmt.Println("Keycard binary downloaded successfully.")
    }

    // Make executable
    err := os.Chmod("./keycard-linux-amd64", 0755)
    if err != nil {
        fmt.Printf("Failed to set execute permission on keycard binary: %s\n", err)
        os.Exit(1)
    }
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
func downloadKeycardBinary() error {
    url := "https://github.com/status-im/keycard-cli/releases/download/0.7.0/keycard-linux-amd64"
    response, err := http.Get(url)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    out, err := os.Create("./keycard-linux-amd64")
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, response.Body)
    return err
}

func checkKeycardBinaryExists() bool {
    if _, err := os.Stat("./keycard-linux-amd64"); os.IsNotExist(err) {
        return false
    }
    return true
}


// generalAskUser asks a general question and returns the user's response.
func generalAskUser(question string) (string, error) {
    fmt.Print(question)
    reader := bufio.NewReader(os.Stdin)
    response, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(response), nil
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

func encryptFile(filename string) error {
    publicKey, err := keycard_link.GetKeycardPublicKey()
    if err != nil {
        return fmt.Errorf("error getting Keycard public key: %w", err)
    }

    passphrase, err := keycard_link.ReadPassphrase()
    if err != nil {
        return fmt.Errorf("error reading passphrase: %w", err)
    }

    seedKDF := publicKey + passphrase
    fmt.Println("SeedKDF:", seedKDF)

    // Derive a key using a KDF (e.g., SHA-256)
    kdfKey := sha256.Sum256([]byte(seedKDF))
    encryptedKey := fmt.Sprintf("%x", kdfKey)

    // Encrypt the file using GPG and the derived key
    cmd := exec.Command("gpg", "--symmetric", "--batch", "--passphrase", encryptedKey, filename)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        return fmt.Errorf("error encrypting file: %w", err)
    }

    fmt.Printf("File encrypted successfully: %s.gpg\n", filename)
    return nil
}

func decryptFile(filename string) error {
    // Ask user for CID
    cid, err := generalAskUser("Enter the CID for the file to decrypt: ")
    if err != nil {
        return fmt.Errorf("error reading CID: %w", err)
    }

    outputPath := "./" // Set the output path, adjust if necessary

    // Retrieve the file from IPFS using ipfs_link library
    err = ipfs_link.GetFileFromIPFS(cid, outputPath)
    if err != nil {
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }

    // Get the Keycard public key
    publicKey, err := keycard_link.GetKeycardPublicKey()
    if err != nil {
        return fmt.Errorf("error getting Keycard public key: %w", err)
    }

    // Read the passphrase
    passphrase, err := keycard_link.ReadPassphrase()
    if err != nil {
        return fmt.Errorf("error reading passphrase: %w", err)
    }

    // Generate the symmetric key
    seedKDF := publicKey + passphrase
    kdfKey := sha256.Sum256([]byte(seedKDF))
    decryptedKey := fmt.Sprintf("%x", kdfKey)

    // Decrypt the file using GPG and the derived key
    decryptedFilePath := outputPath // Adjust as needed
    cmd := exec.Command("gpg", "--decrypt", "--batch", "--passphrase", decryptedKey, "--output", decryptedFilePath, outputPath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        return fmt.Errorf("error decrypting file: %w", err)
    }

    fmt.Printf("File decrypted successfully: %s\n", decryptedFilePath)
    return nil
}




