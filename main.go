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
            fmt.Println("Installing Dependencies...")
            keycard_link.JavaDependency()
            keycard_link.GlobalPlatformDependency()

        case "5":
            fmt.Println("Installing Keycard...")
            err := keycard_link.InstallKeycard()
            if err != nil {
                fmt.Println("Error installing keycard:", err)
            }

	case "6":
	    fmt.Println("Running Connection test to the IPFS Network.")
	    cid := "bafkreie7ohywtosou76tasm7j63yigtzxe7d5zqus4zu3j6oltvgtibeom" // Welcome to IPFS CID
            runIPFSTestWithViu(cid)


        case "7":
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

    // Yellow lines
    fmt.Println("\033[1;33m---------------------------------------------\033[0m")
    // Bold white title
    fmt.Println("\033[1;37mIPFS-Secure | NFC Interface for IPFS\033[0m")
    // Yellow lines
    fmt.Println("\033[1;33m=============================================\033[0m")
    // Normal text for the question
    fmt.Println("   What would you like to do?")
    // Menu options with green arrow, green number, and normal text
    fmt.Println("\033[1;32m>\033[0;32m 1.\033[0m Encrypt / upload sensitive data to IPFS")
    fmt.Println("\033[1;32m>\033[0;32m 2.\033[0m Decrypt / pull file with CID")
    fmt.Println("\033[1;32m>\033[0;32m 3.\033[0m Print CID Log to QR code")
    fmt.Println("\033[1;32m>\033[0;32m 4.\033[0m Install Dependencies (Java, GPP)")
    fmt.Println("\033[1;32m>\033[0;32m 5.\033[0m Install Keycard onto Implant")
    fmt.Println("\033[1;32m>\033[0;32m 6.\033[0m Run Connection Test to IPFS")
    fmt.Println("\033[1;32m>\033[0;32m 7.\033[0m Exit the Program")
    // Yellow lines
    fmt.Println("\033[1;33m=============================================\033[0m")

    return generalAskUser("Enter your choice: ")
}

// runIPFSTestWithViu encapsulates the entire process.
func runIPFSTestWithViu(cid string) {
        if err := checkAndInstallViu(); err != nil {
                fmt.Println("Error installing 'viu':", err)
                fmt.Println("IPFS Check failed tests... :(")
                return
        }

        if err := fetchFromIPFS(cid); err != nil {
                fmt.Println("Error fetching file from IPFS:", err)
                fmt.Println("IPFS Check failed tests... :(")
                return
        }

        if err := displayImage(cid); err != nil {
                fmt.Println("Error displaying image:", err)
                fmt.Println("IPFS Check failed tests... :(")
                return
        }

        if err := performBasicIPFSTests(); err != nil {
                fmt.Println("Error performing basic IPFS tests:", err)
                fmt.Println("IPFS Check failed tests... :(")
                return
        }

        fmt.Println("IPFS tests completed successfully.")
}

// checkAndInstallViu checks if 'viu' is installed and installs it if not.
func checkAndInstallViu() error {
        _, err := exec.LookPath("viu")
        if err != nil {
                fmt.Println("Installing 'viu'...")
                cmd := exec.Command("sudo", "apt-get", "install", "viu", "-y")
                cmd.Stdout = os.Stdout
                cmd.Stderr = os.Stderr
                fmt.Println("\n If this fails, install viu according to your system :) It should be a standard package.. ")
                return cmd.Run()
        }
        return nil
}

func fetchFromIPFS(cid string) error {
    fmt.Println("\033[1;34mFetching from IPFS...\033[0m")
    cmd := exec.Command("ipfs", "get", cid)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func displayImage(filename string) error {
    fmt.Println("\033[1;34mDisplaying image...\033[0m")
    cmd := exec.Command("viu", filename)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func performBasicIPFSTests() error {
    fmt.Println("\033[1;34mPerforming basic IPFS tests...\033[0m")
    os.Remove("bafkreie7ohywtosou76tasm7j63yigtzxe7d5zqus4zu3j6oltvgtibeom")

    // Test 'ipfs diag sys'
    if err := executeIPFSCommand("diag", "sys"); err != nil {
        return fmt.Errorf("\033[1;31mIPFS diag sys test failed: %w\033[0m", err)
    }

    fmt.Println("\033[1;32mAll basic IPFS tests passed successfully.\033[0m")
    return nil
}

func executeIPFSCommand(args ...string) error {
    cmd := exec.Command("ipfs", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
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
        return fmt.Errorf("\033[1;31merror uploading file to IPFS: %w\033[0m", err)
    }

    if err := saveCID(cid); err != nil {
        return fmt.Errorf("\033[1;31merror saving CID: %w\033[0m", err)
    }

    fmt.Printf("\033[1;32mFile uploaded to IPFS successfully, CID:\n%s\033[0m\n", cid)
    fmt.Println("\033[1;33mDeleted standard encrypted format.\033[0m")
    exec.Command("rm", filePath)
    return nil
}

func decryptFileOption() error {
    useLog := askUserYN("\033[1;36mDo you want to use the CID log?\033[0m")

    if useLog {
        return decryptFileFromLog()
    } else {
        filename, err := generalAskUser("\033[1;36mEnter the filename to decrypt ('Save file as'):\033[0m")
        if err != nil {
            return fmt.Errorf("\033[1;31merror reading filename: %w\033[0m", err)
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
        // Question in bold cyan
        fmt.Printf("\033[1;36m%s (y/n):\033[0m ", question)
        response, err := reader.ReadString('\n')
        if err != nil {
            // Error message in bold red
            fmt.Println("\033[1;31mError reading response. Please try again.\033[0m")
            continue
        }
        response = strings.TrimSpace(strings.ToLower(response))

        if response == "y" || response == "yes" {
            return true
        } else if response == "n" || response == "no" {
            return false
        } else {
            // Invalid response in yellow
            fmt.Println("\033[1;33mInvalid response. Please answer 'y' for yes or 'n' for no.\033[0m")
        }
    }
}




func decryptFileFromLog() error {
    cids, err := readCIDLog("log_CID.log")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading CID log: %w\033[0m", err)
    }

    for _, cid := range cids {
        if cid == "" {
            continue
        }

        fmt.Printf("\033[1;34mDecrypting file for CID: %s\033[0m\n", cid)
        err := decryptSingleFile(cid)
        if err != nil {
            fmt.Printf("\033[1;31mError decrypting file for CID %s: %v\033[0m\n", cid, err)
            continue
        }
        fmt.Printf("\033[1;32mFile for CID %s decrypted successfully.\033[0m\n", cid)
    }

    return nil
}



func decryptSingleFile(cid string) error {
    ipfsFilePath := "retrieved_" + cid

    err := ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
    if err != nil {
        return fmt.Errorf("\033[1;31merror retrieving file from IPFS: %w\033[0m", err)
    }

    art_link.PrintFileSlowly("scannow.txt")
    art_link.PrintFileSlowly("flex_implant.txt")

    publicKey, err := keycard_link.GetKeycardPublicKey()
    if err != nil {
        return fmt.Errorf("\033[1;31merror getting Keycard public key: %w\033[0m", err)
    }

    passphrase, err := keycard_link.ReadPassphrase()
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading passphrase: %w\033[0m", err)
    }

    saltPhrase, err := generalAskUser("Enter the salt phrase for decryption: ")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading salt phrase: %w\033[0m", err)
    }

    combinedKey := fmt.Sprintf("%s%s", publicKey, passphrase)

    fmt.Println("\033[1;34mGenerating the key using PBKDF2 for decryption...\033[0m")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New) // 32 byte AES-256

    encryptedData, err := os.ReadFile(ipfsFilePath)
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading encrypted file: %w\033[0m", err)
    }

    decryptedData, err := DecryptAES(encryptedData, aesKey)
    if err != nil {
        fmt.Println("\033[1;33mDid you input the correct passphrase and salt phrase?\033[0m")
        return fmt.Errorf("\033[1;31merror decrypting file: %w\033[0m", err)
    }

    decryptedFilePath := "decrypted_" + cid
    err = os.WriteFile(decryptedFilePath, decryptedData, 0644)
    if err != nil {
        return fmt.Errorf("\033[1;31merror writing decrypted file: %w\033[0m", err)
    }

    fmt.Printf("\033[1;32mFile decrypted successfully: %s\033[0m\n", decryptedFilePath)
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

    publicKey, err := keycard_link.GetKeycardPublicKey()
    if err != nil {
        return fmt.Errorf("\033[1;31merror getting Keycard public key: %w\033[0m", err)
    }

    passphrase, err := keycard_link.ReadPassphrase()
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading passphrase: %w\033[0m", err)
    }
    fmt.Println("\033[1;36mThe machine will generate three random words, write these down (Used in conjunction with your passphrase.)\033[0m")

    saltPhrase, err := generateThreeWordPhrase()
    if err != nil {
        return fmt.Errorf("\033[1;31merror generating salt phrase: %w\033[0m", err)
    }

    fmt.Println("\033[1;32m" + saltPhrase + "\033[0m")
    fmt.Println("\033[1;33mSave these three words and your passphrase to decrypt your files!\033[0m")

    combinedKey := fmt.Sprintf("%s%s", publicKey, passphrase)

    fmt.Println("\033[1;34mGenerating the key using PBKDF2 for encryption...\033[0m")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New) // 32 byte AES-256

    data, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading file to encrypt: %w\033[0m", err)
    }

    encryptedData, err := EncryptAES(data, aesKey)
    if err != nil {
        return fmt.Errorf("\033[1;31merror encrypting file: %w\033[0m", err)
    }

    encryptedFilename := filename + ".aes"
    err = os.WriteFile(encryptedFilename, encryptedData, 0644)
    if err != nil {
        return fmt.Errorf("\033[1;31merror writing encrypted file: %w\033[0m", err)
    }
    fmt.Printf("\033[1;32mFile encrypted successfully: %s\033[0m\n", encryptedFilename)

    if askUserYN("Do you want to upload the encrypted file to IPFS?") {
        if err := ipfsUpload(encryptedFilename); err != nil {
            return fmt.Errorf("\033[1;31merror uploading file to IPFS: %w\033[0m", err)
        }
    }

    return nil
}

func decryptFile(filename string) error {
    cid, err := generalAskUser("Enter the CID for the file to decrypt: ")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading CID: %w\033[0m", err)
    }

    ipfsFilePath := "retrieved_" + filename

    erro := ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
    if erro != nil {
        return fmt.Errorf("\033[1;31merror retrieving file from IPFS: %w\033[0m", erro)
    }

    publicKey, err := keycard_link.GetKeycardPublicKey()
    if err != nil {
        return fmt.Errorf("\033[1;31merror getting Keycard public key: %w\033[0m", err)
    }

    passphrase, err := keycard_link.ReadPassphrase()
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading passphrase: %w\033[0m", err)
    }

    saltPhrase, err := generalAskUser("Enter the salt phrase for decryption: ")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading salt phrase: %w\033[0m", err)
    }

    combinedKey := fmt.Sprintf("%s%s", publicKey, passphrase)

    fmt.Println("\033[1;34mGenerating the key using PBKDF2 for decryption...\033[0m")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New) // 32 byte AES-256

    encryptedData, err := os.ReadFile(ipfsFilePath)
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading encrypted file: %w\033[0m", err)
    }

    decryptedData, err := DecryptAES(encryptedData, aesKey)
    if err != nil {
        return fmt.Errorf("\033[1;31merror decrypting file: %w\033[0m", err)
    }

    decryptedFilePath := "decrypted_" + cid

    err = os.WriteFile(decryptedFilePath, decryptedData, 0644)
    if err != nil {
        return fmt.Errorf("\033[1;31merror writing decrypted file: %w\033[0m", err)
    }

    fmt.Printf("\033[1;32mFile decrypted successfully: %s\033[0m\n", decryptedFilePath)
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



