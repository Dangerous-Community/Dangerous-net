package ipfs_link

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
    "Dangerous-net/art_link"
)

// AddFileToIPFS adds a file to IPFS and returns the CID (Content Identifier)
func AddFileToIPFS(filePath string) (string, error) {
    done := make(chan bool)
    go art_link.LoadingScreen(done) // Start the loading screen in a separate goroutine

    cmd := exec.Command("ipfs", "add", filePath)

    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        done <- true // Signal to stop the loading screen on error
        return "", fmt.Errorf("error adding file to IPFS: %w", err)
    }

    output := out.String()
    // Extract CID from the output
    lines := strings.Split(output, "\n")
    for _, line := range lines {
        if strings.Contains(line, "added") {
            parts := strings.Fields(line)
            if len(parts) >= 2 {
                done <- true // Signal to stop the loading screen on successful add
                return parts[1], nil // Assuming the CID is the second part
            }
        }
    }

    done <- true // Signal to stop the loading screen if CID is not found
    return "", fmt.Errorf("CID not found in IPFS add output")
}
// GetFileFromIPFS retrieves a file from IPFS using its CID
func GetFileFromIPFS(cid, outputPath string) error {
    done := make(chan bool)
    go art_link.LoadingScreen(done) // Start the loading screen in a separate goroutine

    cmd := exec.Command("ipfs", "get", cid, "-o", outputPath)
    err := cmd.Run()
    if err != nil {
        done <- true // Signal to stop the loading screen on error
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }

    done <- true // Signal to stop the loading screen on successful download
    return nil
}
