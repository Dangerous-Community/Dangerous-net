package chat_dapp

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "bytes"
)

// chatDappScriptURL is the URL of the chat dapp script.
const chatDappScriptURL = "https://raw.githubusercontent.com/SomajitDey/ipfs-chat/main/ipfs-chat"

// GetScriptPath determines the path of the script.
func GetScriptPath() (string, error) {
    exePath, err := os.Executable()
    if err != nil {
        return "", err
    }
    dir := filepath.Dir(exePath)
    return filepath.Join(dir, "ipfs-chat"), nil
}

// DownloadScript downloads the script from the specified URL to the given path.
func DownloadScript(path string) error {
    resp, err := http.Get(chatDappScriptURL)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    outFile, err := os.Create(path)
    if err != nil {
        return err
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, resp.Body)
    if err != nil {
        return err
    }

    // Setting file permissions to 750
    if err := os.Chmod(path, 0750); err != nil {
        return fmt.Errorf("failed to set file permissions: %w", err)
    }

    return nil
}

// InstallBashScript checks if the script is available and downloads it if not.
func InstallBashScript() error {
    scriptPath, err := GetScriptPath()
    if err != nil {
        fmt.Printf("Error obtaining script path: %s\n", err)
        return err
    }

    if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
        fmt.Println("Script not found, downloading...")
        if err := DownloadScript(scriptPath); err != nil {
            fmt.Printf("Error downloading script: %s\n", err)
            return err
        }
        fmt.Println("Script downloaded successfully.")
    } else {
        fmt.Println("Script already exists.")
    }

    return nil
}

// RunBashScript executes the chat dapp bash script.
func RunBashScript() error {
    scriptPath, err := GetScriptPath()
    if err != nil {
        fmt.Printf("Error obtaining script path: %s\n", err)
        return err
    }

    cmd := exec.Command("bash", scriptPath)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error executing script:", err)
        fmt.Println("Script output:", stdout.String())
        fmt.Println("Script error:", stderr.String())
        fmt.Println("If the script failed due to missing dependencies, please install 'jq' and 'dialog' using your package manager.")
        return err
    }

    fmt.Println("Script executed successfully:", stdout.String())
    return nil
}
