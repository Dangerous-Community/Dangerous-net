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
    "Dangerous-net/keycard_link"
    "github.com/mdp/qrterminal/v3"
    "Dangerous-net/art_link"
    "Dangerous-net/ipfs_link"
    "Dangerous-net/chat_dapp"
)

//go:embed english.txt
var englishTxt embed.FS


func main() {
    // Setup configuration file
    if err := setupConfig(); err != nil {
        fmt.Printf("Failed to set up config: %s\n", err)
        os.Exit(1)
    }



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
        if handleFileManagement() {
            continue // If true, continue the main loop, effectively going back to main menu
	}
    case "2":
        if err := chat_dapp.RunBashScript(); err != nil {
            fmt.Println("Error running chat DApp:", err)
        }
    case "3":
        qr() // Generate QR code
    case "4":
        fmt.Println("Installing Dependencies...")
        keycard_link.JavaDependency()
        keycard_link.GlobalPlatformDependency()
    case "5":
        fmt.Println("Installing Keycard...")
        if err := keycard_link.InstallKeycard(); err != nil {
            fmt.Println("Error installing keycard:", err)
        }
    case "6":
        fmt.Println("Running Connection test to the IPFS Network.")
        runIPFSTestWithViu("bafkreie7ohywtosou76tasm7j63yigtzxe7d5zqus4zu3j6oltvgtibeom") // CID for test
    case "7":
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

    fmt.Println("\033[1;32m>\033[0;32m 1.\033[0m File Management")
    fmt.Println("\033[1;32m>\033[0;32m 2.\033[0m Run Chat DApp")
    fmt.Println("\033[1;32m>\033[0;32m 3.\033[0m Print CID Log to QR code")
    fmt.Println("\033[1;32m>\033[0;32m 4.\033[0m Install Dependencies (Java, GPP)")
    fmt.Println("\033[1;32m>\033[0;32m 5.\033[0m Install Keycard onto Implant")
    fmt.Println("\033[1;32m>\033[0;32m 6.\033[0m Run Connection Test to IPFS")
    fmt.Println("\033[1;32m>\033[0;32m 7.\033[0m Exit the Program")
    // Yellow lines
    // Yellow lines
    fmt.Println("\033[1;33m=============================================\033[0m")

    return generalAskUser("Enter your choice: ")
}

func handleFileManagement()bool {
    for {
        choice, err := fileManagementMenu()
        if err != nil {
            fmt.Println("Error:", err)
            return false
        }
        switch choice {
        case "1":
            filename, err := generalAskUser("Enter the filename to encrypt and upload: ")
            if err != nil {
                fmt.Println("Error:", err)
                continue
            }
            if err := encryptAndUploadFile(filename); err != nil {
                fmt.Println("Error:", err)
            }

        case "2":
            if err := decryptAndDownloadFile(); err != nil {
                fmt.Println("Error:", err)
            }

        case "3":
            savePath, err := ipfsDownload()
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Printf("File downloaded successfully to %s\n", savePath)
            }

        case "4":
            filePath, err := generalAskUser("Enter the file path to upload to IPFS: ")
            if err != nil {
                fmt.Println("Error:", err)
                continue
            }
            level, err := readConfig()
            if err != nil {
                fmt.Println("Error:", err)
                continue
            }
            if err := ipfsUpload(filePath, level); err != nil {
                fmt.Println("Error:", err)
            }
	case "5":
	     return true

        default:
            fmt.Println("Invalid option, please try again.")
        }
    }
}

func fileManagementMenu() (string, error) {
    err := art_link.PrintFileSlowly("ipfs.txt")
    if err != nil {
        fmt.Println("Error displaying ASCII art:", err)
    }

    // Yellow lines
    fmt.Println("\033[1;33m---------------------------------------------\033[0m")
    // Bold white title
    fmt.Println("\033[1;37mDangerous Net | File Management Menu\033[0m")
    // Yellow lines
    fmt.Println("\033[1;33m=============================================\033[0m")
    // Normal text for the question
    fmt.Println("   What would you like to do?")

    fmt.Println("\033[1;32m>\033[0;32m 1.\033[0m Encrypt and Upload")
    fmt.Println("\033[1;32m>\033[0;32m 2.\033[0m Decrypt and Download")
    fmt.Println("\033[1;32m>\033[0;32m 3.\033[0m Just Download from IPFS")
    fmt.Println("\033[1;32m>\033[0;32m 4.\033[0m Just Upload to IPFS")
    fmt.Println("\033[1;32m>\033[0;32m 5.\033[0m Return to Main Menu")
    // Yellow lines
    fmt.Println("\033[1;33m=============================================\033[0m")

    return generalAskUser("Enter your choice: ")
}

func encryptAndUploadFile(filename string) error {
    if err := encryptFile(filename); err != nil {
        return err
    }
    encryptedFilename := filename + ".aes"
    level, err := readConfig()
    if err != nil {
        return err
    }
    return ipfsUpload(encryptedFilename, level)
}


func decryptAndDownloadFile() error {
//    IPFS DOWNLOAD IS ALREADY CALLED WITHIN DECRYPT FILE FUNCTION! :)
	return decryptFile()
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




func saveCID(cid string, level string) error {
    f, err := os.OpenFile("log_CID.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    logEntry := fmt.Sprintf("%s,%s\n", cid, level) // Format: "cid,level"
    if _, err := f.WriteString(logEntry); err != nil {
        return err
    }

    return nil
}



func decryptFileOption() error {
    useLog := askUserYN("\033[1;36mDo you want to use the CID log?\033[0m")

    if useLog {
        return decryptFileFromLog()
    } else {
        return decryptFile()
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



func getEncryptionLevelForCID(cid string) (string, error) {
    logFile := os.Getenv("HOME") + "/.config/DangerousNet/log_CID.log"
    data, err := os.ReadFile(logFile)
    if err != nil {
        return "", fmt.Errorf("error reading CID log: %w", err)
    }

    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        parts := strings.Split(line, ",")
        if len(parts) == 2 && parts[0] == cid {
            return parts[1], nil
        }
    }

    // CID not found in log, ask the user to select the level manually
    fmt.Println("Sorry, the level could not be found in the log. Which one is it for this file?")
    fmt.Println("  > 1 Easy")
    fmt.Println("  > 2 Medium")
    fmt.Println("  > 3 Hard")
    reader := bufio.NewReader(os.Stdin)
    choice, err := reader.ReadString('\n')
    if err != nil {
        return "", fmt.Errorf("error reading user input: %w", err)
    }

    choice = strings.TrimSpace(choice)
    switch choice {
    case "1":
        return "easy", nil
    case "2":
        return "medium", nil
    case "3":
        return "hard", nil
    default:
        return "", fmt.Errorf("invalid selection")
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

        level, err := getEncryptionLevelForCID(cid) // Correctly handle both returned values
        if err != nil {
            fmt.Printf("\033[1;31mError retrieving encryption level for CID %s: %v\033[0m\n", cid, err)
            continue
        }

        fmt.Printf("\033[1;34mDecrypting file for CID: %s\033[0m\n", cid)
        err = decryptSingleFile(cid, level)
        if err != nil {
            fmt.Printf("\033[1;31mError decrypting file for CID %s: %v\033[0m\n", cid, err)
            continue
        }
        fmt.Printf("\033[1;32mFile for CID %s decrypted successfully.\033[0m\n", cid)
    }

    return nil
}

func decryptSingleFile(cid string, level string) error {
    ipfsFilePath := "retrieved_" + cid

    err := ipfs_link.GetFileFromIPFS(cid, ipfsFilePath)
    if err != nil {
        return fmt.Errorf("\033[1;31merror retrieving file from IPFS: %w\033[0m", err)
    }

     var combinedKey string
    var saltPhrase string
    var passphrase string

    if level == "medium" || level == "hard" {
        passphrase, err = generalAskUser("Enter your passphrase for the file: ")
        if err != nil {
            return fmt.Errorf("\033[1;31merror reading passphrase: %w\033[0m", err)
        }
    }

    saltPhrase, err = generalAskUser("Enter the salt phrase used for encryption (Machine generated 3 words): ")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading salt phrase: %w\033[0m", err)
    }

    if level == "hard" {
        publicKey, err := keycard_link.GetKeycardPublicKey()
        if err != nil {
            return fmt.Errorf("\033[1;31merror getting Keycard public key: %w\033[0m", err)
        }
        combinedKey = passphrase + saltPhrase + publicKey
    } else {
        combinedKey = passphrase + saltPhrase
    }

    fmt.Println("\033[1;34mGenerating the key using PBKDF2 for decryption...\033[0m")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New)

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


func ipfsUpload(filePath string, level string) error {
    cid, err := ipfs_link.AddFileToIPFS(filePath)
    if err != nil {
        return fmt.Errorf("\033[1;31merror uploading file to IPFS: %w\033[0m", err)
    }

    if err := saveCID(cid, level); err != nil {
        return fmt.Errorf("\033[1;31merror saving CID and level: %w\033[0m", err)
    }

    fmt.Printf("\033[1;32mFile uploaded to IPFS successfully, CID:\n%s\033[0m\n", cid)
    fmt.Println("\033[1;33mDeleted standard encrypted format.\033[0m")
    exec.Command("rm", filePath)
    return nil
}



func ipfsDownload() (string, error) {
    cid, err := generalAskUser("Enter the CID of the file to download from IPFS: ")
    if err != nil {
        return "", fmt.Errorf("\033[1;31merror reading CID: %w\033[0m", err)
    }

    savePath, err := generalAskUser("Enter the path to save the downloaded file: ")
    if err != nil {
        return "", fmt.Errorf("\033[1;31merror reading save path: %w\033[0m", err)
    }

    err = ipfs_link.GetFileFromIPFS(cid, savePath)
    if err != nil {
        return "", fmt.Errorf("\033[1;31merror downloading file from IPFS: %w\033[0m", err)
    }

    fmt.Printf("\033[1;32mFile downloaded successfully: %s\033[0m\n", savePath)
    return savePath, nil
}

func decryptFile() error {
    filePath, err := generalAskUser("Enter the path of the encrypted file to decrypt: ")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading file path: %w\033[0m", err)
    }

    level, err := generalAskUser("Enter the encryption level used on the file (easy, medium, hard): ")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading encryption level: %w\033[0m", err)
    }

    var combinedKey, passphrase, saltPhrase string
    if level == "medium" || level == "hard" {
        passphrase, err = generalAskUser("Enter your passphrase for the file: ")
        if err != nil {
            return fmt.Errorf("\033[1;31merror reading passphrase: %w\033[0m", err)
        }
    }

    saltPhrase, err = generalAskUser("Enter the salt phrase used for encryption (Machine generated 3 words): ")
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading salt phrase: %w\033[0m", err)
    }

    if level == "hard" {
        publicKey, err := keycard_link.GetKeycardPublicKey()
        if err != nil {
            return fmt.Errorf("\033[1;31merror getting Keycard public key: %w\033[0m", err)
        }
        combinedKey = passphrase + saltPhrase + publicKey
    } else {
        combinedKey = passphrase + saltPhrase
    }

    fmt.Println("\033[1;34mGenerating the key using PBKDF2 for decryption...\033[0m")
    combinedKeyBytes := []byte(combinedKey)
    saltBytes := []byte(saltPhrase)
    iterationCount := 1000000
    aesKey := pbkdf2.Key(combinedKeyBytes, saltBytes, iterationCount, 32, sha256.New)

    encryptedData, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("\033[1;31merror reading encrypted file: %w\033[0m", err)
    }

    decryptedData, err := DecryptAES(encryptedData, aesKey)
    if err != nil {
        fmt.Println("\033[1;33mDid you input the correct passphrase and salt phrase?\033[0m")
        return fmt.Errorf("\033[1;31merror decrypting file: %w\033[0m", err)
    }

    decryptedFilePath := "decrypted_" + filePath
    err = os.WriteFile(decryptedFilePath, decryptedData, 0644)
    if err != nil {
        return fmt.Errorf("\033[1;31merror writing decrypted file: %w\033[0m", err)
    }

    fmt.Printf("\033[1;32mFile decrypted successfully: %s\033[0m\n", decryptedFilePath)
    return nil
}


func setupConfig() error {
    configDir := os.Getenv("HOME") + "/.config/DangerousNet"
    configFile := configDir + "/config"

    // Create the DangerousNet directory if it doesn't exist
    err := os.MkdirAll(configDir, 0755)
    if err != nil {
        return fmt.Errorf("error creating config directory: %w", err)
    }

    // Check if the config file exists
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        // Create a default config file
        defaultConfig := []byte("encryptionLevel=easy\n")
        err = os.WriteFile(configFile, defaultConfig, 0644)
        if err != nil {
            return fmt.Errorf("error creating default config file: %w", err)
        }
    }

    return nil
}

func readConfig() (string, error) {
    configFile := os.Getenv("HOME") + "/.config/DangerousNet/config"
    data, err := os.ReadFile(configFile)
    if err != nil {
        return "", fmt.Errorf("error reading config file: %w", err)
    }

    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        if strings.HasPrefix(line, "encryptionLevel=") {
            return strings.TrimPrefix(line, "encryptionLevel="), nil
        }
    }

    return "medium", nil // Default level if not specified in the config
}


func encryptFile(filename string ) error {
    level, erro := readConfig()
    if erro != nil {
        return fmt.Errorf("error reading encryption level from config: %w", erro)
    }
    var passphrase string
    var err error

    if level == "medium" || level == "hard" {
        passphrase, err = generalAskUser("Enter a passphrase for this file: ")
        if err != nil {
            return fmt.Errorf("\033[1;31merror reading passphrase: %w\033[0m", err)
        }
    }

    saltPhrase, err := generateThreeWordPhrase()
    if err != nil {
        return fmt.Errorf("\033[1;31merror generating salt phrase: %w\033[0m", err)
    }

    var combinedKey string
    if level == "hard" {
        publicKey, err := keycard_link.GetKeycardPublicKey()
        if err != nil {
            return fmt.Errorf("\033[1;31merror getting Keycard public key: %w\033[0m", err)
        }
        combinedKey = passphrase + publicKey // Using passphrase and publicKey for hard level
    } else if level == "medium" {
        combinedKey = passphrase // Using only passphrase for medium level
    } else {
        combinedKey = "" // No passphrase for easy level
    }

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
        if err := ipfsUpload(encryptedFilename, level); err != nil {
            return fmt.Errorf("\033[1;31merror uploading file to IPFS: %w\033[0m", err)
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



