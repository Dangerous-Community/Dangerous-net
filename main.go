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
    "golang.org/x/crypto/pbkdf2"
    "crypto/rand"
    "crypto/cipher"
    "crypto/aes"
    "embed"
    "math/big"
    "IPFSS_IPFS-Secure/keycard_link"
    "github.com/mdp/qrterminal/v3"
    "IPFSS_IPFS-Secure/art_link"
    "IPFSS_IPFS-Secure/ipfs_link"
)

//go:embed english.txt
var englishTxt embed.FS


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

    err := os.Chmod("./keycard-linux-amd64", 0755)
    if err != nil {
        fmt.Printf("Failed to set execute permission on keycard binary: %s\n", err)
        os.Exit(1)
    }

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
            qr()

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




func generalAskUser(question string) (string, error) {
    fmt.Print(question)
    reader := bufio.NewReader(os.Stdin)
    response, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(response), nil
}




func saveCID(cid string) error {
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
    cid, err := ipfs_link.AddFileToIPFS(filePath)
    if err != nil {
        return err
    }

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
    cids, err := readCIDLog("log_CID.log")
    if err != nil {
        return fmt.Errorf("error reading CID log: %w", err)
    }

    for _, cid := range cids {
        if cid == "" {
            continue
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
    cid, err := generalAskUser("Enter the CID for the file to decrypt: ")
    if err != nil {
        return fmt.Errorf("error reading CID: %w", err)
    }

    ipfsFilePath := "retrieved_" + filename

    erro := ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
    if erro != nil {
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }

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



    saltPhrase, err := generalAskUser("Enter the salt phrase for decryption: ")
    if err != nil {
        return fmt.Errorf("error reading salt phrase: %w", err)
    }

    combinedKey := fmt.Sprintf("%s%s", publicKey, passphrase)

    fmt.Println("Generating the key using PBKDF2 for decryption...")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New) // 32 byte AES-256

    encryptedData, err := os.ReadFile(ipfsFilePath)
    if err != nil {
        return fmt.Errorf("error reading encrypted file: %w", err)
    }

    decryptedData, err := DecryptAES(encryptedData, aesKey)
    if err != nil {
        return fmt.Errorf("error decrypting file: %w", err)
    }

    decryptedFilePath := "decrypted_" + cid

    err = os.WriteFile(decryptedFilePath, decryptedData, 0644)
    if err != nil {
        return fmt.Errorf("error writing decrypted file: %w", err)
    }

    fmt.Printf("File decrypted successfully: %s\n", decryptedFilePath)
    return nil
}




func decryptSingleFile(cid string) error {
    ipfsFilePath := "retrieved_" + cid

    err := ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
    if err != nil {
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }


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

    saltPhrase, err := generalAskUser("Enter the salt phrase for decryption: ")
    if err != nil {
        return fmt.Errorf("error reading salt phrase: %w", err)
    }

    combinedKey := fmt.Sprintf("%s%s", publicKey, passphrase)

    fmt.Println("Generating the key using PBKDF2 for decryption...")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New) // 32 byte AES-256

    encryptedData, err := os.ReadFile(ipfsFilePath)
    if err != nil {
        return fmt.Errorf("error reading encrypted file: %w", err)
    }

    decryptedData, err := DecryptAES(encryptedData, aesKey)
    if err != nil {
        fmt.Println("Did you input the correct passphrase and salt phrase?")
        return fmt.Errorf("error decrypting file: %w", err)
    }

    decryptedFilePath := "decrypted_" + cid
    err = os.WriteFile(decryptedFilePath, decryptedData, 0644)

    if err != nil {
        return fmt.Errorf("error writing decrypted file: %w", err)
    }

    fmt.Printf("File decrypted successfully: %s\n", decryptedFilePath)
    return nil
}

func generateThreeWordPhrase() (string, error) {
    content, err := englishTxt.ReadFile("english.txt")
    if err != nil {
        return "", err
    }

    dictionary := strings.Split(string(content), "\n")
    var phrase []string
    for i := 0; i < 3; i++ {
        idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(dictionary))))
        if err != nil {
            return "", err
        }
        phrase = append(phrase, dictionary[idx.Int64()])
    }
    return fmt.Sprintf("%s %s %s", phrase[0], phrase[1], phrase[2]), nil
}

func encryptFile(filename string) error {
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
    fmt.Println("The machine will generate three random words, write these down (Used in conjunciton with your passphrase.)")

    saltPhrase, err := generateThreeWordPhrase()
    if err != nil {
        return fmt.Errorf("error generating salt phrase: %w", err)
    }

    fmt.Println(saltPhrase)
    fmt.Println("Save these three words and your passphrase to decrypt your files!")


    combinedKey := fmt.Sprintf("%s%s", publicKey, passphrase)

    fmt.Println("Generating the key using PBKDF2 for encryption...")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New) // 32 byte AES-256

    data, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("error reading file to encrypt: %w", err)
    }

    encryptedData, err := EncryptAES(data, aesKey)
    if err != nil {
        return fmt.Errorf("error encrypting file: %w", err)
    }

    encryptedFilename := filename + ".aes"
    err = os.WriteFile(encryptedFilename, encryptedData, 0644)
    if err != nil {
        return fmt.Errorf("error writing encrypted file: %w", err)
    }
    fmt.Printf("File encrypted successfully: %s\n", encryptedFilename)

    if askUserYN("Do you want to upload the encrypted file to IPFS?") {
        if err := ipfsUpload(encryptedFilename); err != nil {
            return fmt.Errorf("error uploading file to IPFS: %w", err)
        }
    }

    return nil
}

func EncryptAES(data []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    return gcm.Seal(nonce, nonce, data, nil), nil
}

func DecryptAES(data []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return nil, err
    }

    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}



