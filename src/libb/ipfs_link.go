package ipfs_link

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
)

// AddFileToIPFS adds a file to IPFS and returns the CID (Content Identifier)
func AddFileToIPFS(filePath string) (string, error) {
    cmd := exec.Command("ipfs", "add", filePath)

    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("error adding file to IPFS: %w", err)
    }

    output := out.String()
    // Extract CID from the output
    lines := strings.Split(output, "\n")
    if len(lines) < 2 {
        return "", fmt.Errorf("unexpected output from IPFS add")
    }
    parts := strings.Fields(lines[1])
    if len(parts) < 2 {
        return "", fmt.Errorf("unexpected output format from IPFS add")
    }
    cid := parts[1]

    return cid, nil
}

// GetFileFromIPFS retrieves a file from IPFS using its CID
func GetFileFromIPFS(cid, outputPath string) error {
    cmd := exec.Command("ipfs", "get", cid, "-o", outputPath)

    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("error retrieving file from IPFS: %w", err)
    }

    return nil
}
