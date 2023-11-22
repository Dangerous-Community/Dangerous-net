package main

import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "bufio"
    "net/http"
    "crypto/sha256"
    "encoding/hex"
    "strings"
    "IPFSS_IPFS-Secure/keycard_link"
    "IPFSS_IPFS-Secure/art_link"
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

    for {
        choice, err := menu()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        switch choice {
        case "1":
            filename, err := generalAskUser("Enter the filename to decrypt ('Save file as'): ")
            if err != nil {
                fmt.Println("Error:", err)
                continue
            }
            if err := decryptFile(filename); err != nil {
                fmt.Println("Error:", err)
            }

        case "2":
            filename, err := generalAskUser("Enter the filename to encrypt: ")
            if err != nil {
                fmt.Println("Error:", err)
                continue
            }
            if err := encryptFile(filename); err != nil {
                fmt.Println("Error:", err)
            }

        case "3":
            err := art_link.PrintFileSlowly("apexflexflexsecure.txt")
            if err != nil {
                fmt.Println("Error displaying ASCII art:", err)
            }

            fmt.Println("Exiting...")
            os.Exit(0)

        default:
            fmt.Println("Invalid option, please try again.")
        }
    }
}

func menu() (string, error) {
    err := art_link.PrintFileSlowly("ipfs.txt")
    if err != nil {
        fmt.Println("Error displaying ASCII art:", err)
    }

    fmt.Println("=============================================")
    fmt.Println("What would you like to do? Select 1, 2, or 3")
    fmt.Println("1. Decrypt / pull file with CID.")
    fmt.Println("2. Encrypt / upload sensitive data to IPFS.")
    fmt.Println("3. Exit.")
    fmt.Println("=============================================")
    return generalAskUser("Enter your choice: ")
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
    // Get the Keycard public key
    art_link.PrintFileSlowly("scannow.txt")
    art_link.PrintFileSlowly("flex_implant.txt")

    publicKey, err := keycard_link.GetKeycardPublicKey()
    if err != nil {
        return fmt.Errorf("error getting Keycard public key: %w", err)
    }

    passphrase, err := keycard_link.ReadPassphrase()
    if err != nil {
        return fmt.Errorf("error reading passphrase: %w", err)
    }
    fmt.Print("Generating the seed for KDF ... ")
    // Get the public key from the server
    publicKey, _ := keys.NewECDSAKeyFromPEM(nil)
    // Convert the public key to a byte slice
    pubKeyBytes := []byte(publicKey)
    // Convert the passphrase to a byte slice
    passphraseBytes := []byte(passphrase)
    // Concatenate the two byte slices
    seedKDF := append(pubKeyBytes[:], passphraseBytes...)
    // Derive a key using a KDF (e.g., SHA-256)
    kdfKey := sha256.Sum256(seedKDF)
    fmt.Println("KDF For symmetric keygen: ", kdfKey)
    fmt.Print("Generating the symmetric key... \n")
    encryptedKey := hex.EncodeToString(kdfKey[:])
    art_link.PrintFileSlowly("encrypting.txt")
    

    // Encrypt the file using GPG and the derived key
    cmd := exec.Command("gpg", "--symmetric", "--batch", "--passphrase", encryptedKey, filename)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        return fmt.Errorf("error encrypting file: %w", err)
    }
    fmt.Printf("File encrypted successfully: %s.gpg\n", filename)

    // Use askUserYN to ask if the user wants to upload to IPFS
    if askUserYN("Do you want to upload the encrypted file to IPFS?") {
        // Call ipfsUpload function with the encrypted file
        if err := ipfsUpload(filename + ".gpg"); err != nil {
            return fmt.Errorf("error uploading file to IPFS: %w", err)
        }
    }

    return nil
}

func saveCID(cid string) error {
    // Write CID to log_CID.log
    f, err := os.OpenFile("log_CID.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    if _, err := f.WriteString(cid + "\n"); err != nil {
        return err
    }

    return nil
}


func ipfsUpload(filePath string) error {
    // Call AddFileToIPFS function
    cid, err := ipfs_link.AddFileToIPFS(filePath)
    if err != nil {
        return err
    }

    // Call saveCID function with the returned CID
    if err := saveCID(cid); err != nil {
        return fmt.Errorf("error saving CID: %w", err)
    }

    fmt.Println("File uploaded to IPFS successfully, CID:", cid)
    return nil
}


func decryptFile(filename string) error {
    // Ask user for CID
    cid, err := generalAskUser("Enter the CID for the file to decrypt: ")
    if err != nil {
        return fmt.Errorf("error reading CID: %w", err)
    }

    ipfsFilePath := "retrieved_" + filename // Save the file retrieved from IPFS with a prefixed name

    // Retrieve the file from IPFS
    err = ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
    if err != nil {
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }

    // Get the Keycard public key
    art_link.PrintFileSlowly("scannow.txt")
    art_link.PrintFileSlowly("flex_implant.txt")

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
    fmt.Print("Generating the symmetric key... \n")
    seedKDF := publicKey + passphrase
    kdfKey := sha256.Sum256([]byte(seedKDF))
    decryptedKey := fmt.Sprintf("%x", kdfKey)
    art_link.PrintFileSlowly("decrypting.txt")

    // Decrypt the file using GPG
    decryptedFilePath := "decrypted_" + filename // This is the path where the decrypted file will be saved
    cmd := exec.Command("gpg", "--decrypt", "--batch", "--passphrase", decryptedKey, "--output", decryptedFilePath, ipfsFilePath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        return fmt.Errorf("error decrypting file: %w", err)
    }

    fmt.Printf("File decrypted successfully: %s\n", decryptedFilePath)
    return nil
}

