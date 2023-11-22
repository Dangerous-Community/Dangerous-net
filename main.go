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
    "github.com/mdp/qrterminal/v3"
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
            filename, err := generalAskUser("Enter the filename to encrypt: ")
            if err != nil {
                fmt.Println("Error:", err)
                continue
            }
            if err := encryptFile(filename); err != nil {
                fmt.Println("Error:", err)
            }

        case "2":
            err := decryptFileOption()
            if err != nil {
                fmt.Println("Error:", err)
            }

        case "3":
            qr() // Call the function to print the QR code

        case "4":
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

func qr(){

    // Read the file
    data, err := os.ReadFile("log_CID.log")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    dataStr := string(data)

    config := qrterminal.Config{
        Level: qrterminal.L,
        Writer: os.Stdout,
        BlackChar: qrterminal.BLACK,
        WhiteChar: qrterminal.WHITE,
        QuietZone: 1,
    }

    qrterminal.GenerateWithConfig(dataStr, config)

    fmt.Println("QR code generated and printed successfully.")

}


func menu() (string, error) {
    err := art_link.PrintFileSlowly("ipfs.txt")
    if err != nil {
        fmt.Println("Error displaying ASCII art:", err)
    }
    fmt.Println("---------------------------------------------")
    fmt.Println("IPFS-Secure | NFC Interface for IPFS ")
    fmt.Println("=============================================")
    fmt.Println("What would you like to do? Select 1, 2, or 3")
    fmt.Println("1. Encrypt / upload sensitive data to IPFS.")
    fmt.Println("2. Decrypt / pull file with CID.")
    fmt.Println("3. Print CID Log to QR code.")
    fmt.Println("4. Exit.")
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


func decryptSingleFile(cid string) error {
    ipfsFilePath := "retrieved_" + cid // Adjust filename as needed

    // Retrieve the file from IPFS
    err := ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
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

    passphrase, err := keycard_link.ReadPassphrase()
    if err != nil {
        return fmt.Errorf("error reading passphrase: %w", err)
    }

    fmt.Print("Generating the seed for KDF ... ")
    // Convert the public key to a byte slice
    pubKeyBytes := []byte(publicKey)
    // Convert the passphrase to a byte slice
    passphraseBytes := []byte(passphrase)
    // Concatenate the two byte slices
    seedKDF := append(pubKeyBytes[:], passphraseBytes...)
    // Derive a key using a KDF (e.g., SHA-256)
    kdfKey := sha256.Sum256(seedKDF)
    fmt.Println("KDF For symmetric keygen: \n", kdfKey)
    fmt.Print("Generating the symmetric key... \n")
    decryptedKey := hex.EncodeToString(kdfKey[:])
    art_link.PrintFileSlowly("decrypting.txt")


    // Decrypt the file using GPG
    decryptedFilePath := "decrypted_" + cid // This is the path where the decrypted file will be saved
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
    // Convert the public key to a byte slice
    pubKeyBytes := []byte(publicKey)
    // Convert the passphrase to a byte slice
    passphraseBytes := []byte(passphrase)
    // Concatenate the two byte slices
    seedKDF := append(pubKeyBytes[:], passphraseBytes...)
    // Derive a key using a KDF (e.g., SHA-256)
    kdfKey := sha256.Sum256(seedKDF)
    fmt.Println("KDF For symmetric keygen: \n", kdfKey)
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

    fmt.Println("File uploaded to IPFS successfully, CID:\n", cid)
    fmt.Println("Deleted standard encrypted format.")
    exec.Command("rm" , filePath)
    return nil
}

func decryptFileOption() error {
    useLog := askUserYN("Do you want to use the CID log?")

    if useLog {
        return decryptFileFromLog()
    } else {
        filename, err := generalAskUser("Enter the filename to decrypt ('Save file as'): ")
        if err != nil {
            return err
        }
        return decryptFile(filename)
    }
}

func readCIDLog(logFile string) ([]string, error) {
    data, err := os.ReadFile(logFile)
    if err != nil {
        return nil, err
    }

    // Assuming each line in the log file is a separate CID
    lines := strings.Split(string(data), "\n")
    return lines, nil
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




func decryptFileFromLog() error {
    // Read CID log
    cids, err := readCIDLog("log_CID.log")
    if err != nil {
        return fmt.Errorf("error reading CID log: %w", err)
    }

    // Loop through each CID in the log and decrypt the associated file
    for _, cid := range cids {
        if cid == "" {
            continue // Skip empty lines
        }

        fmt.Printf("Decrypting file for CID: %s\n", cid)
        err := decryptSingleFile(cid)
        if err != nil {
            fmt.Printf("Error decrypting file for CID %s: %v\n", cid, err)
            continue
        }
        fmt.Printf("File for CID %s decrypted successfully.\n", cid)
    }

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
    erro := ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
    if erro != nil {
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }

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
    // Convert the public key to a byte slice
    pubKeyBytes := []byte(publicKey)
    // Convert the passphrase to a byte slice
    passphraseBytes := []byte(passphrase)
    // Concatenate the two byte slices
    seedKDF := append(pubKeyBytes[:], passphraseBytes...)
    // Derive a key using a KDF (e.g., SHA-256)
    kdfKey := sha256.Sum256(seedKDF)
    fmt.Println("KDF For symmetric keygen: \n", kdfKey)
    fmt.Print("Generating the symmetric key... \n")
    decryptedKey := hex.EncodeToString(kdfKey[:])
    art_link.PrintFileSlowly("decrypting.txt")


    // Decrypt the file using GPG
    decryptedFilePath := "decrypted_" + cid // Use cid instead of filename
    cmd := exec.Command("gpg", "--decrypt", "--batch", "--passphrase", decryptedKey, "--output", decryptedFilePath, ipfsFilePath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Run()
    if err != nil {
        return fmt.Errorf("error decrypting file: %w", err)
    }

    fmt.Printf("File decrypted successfully: %s\n", decryptedFilePath)
    return nil // Remove the extra return statement
}
