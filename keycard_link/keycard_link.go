package keycard_link

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// GetKeycardPublicKey retrieves the public key from the keycard.
func GetKeycardPublicKey() (string, error) {
    // Command to execute
    fmt.Print("Scan your card now! :) \n")
    cmd := exec.Command("./keycard-linux-amd64", "info")

    // Capture the output of the command
    var out bytes.Buffer
    cmd.Stdout = &out

    // Run the command
    err := cmd.Run()
    if err != nil {
        return "", err
    }

    // Process the output to find the public key
    scanner := bufio.NewScanner(&out)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "PublicKey:") {
            // Assuming the public key is the last element in the line, separated by spaces
            parts := strings.Fields(line)
            if len(parts) > 1 {
                return parts[len(parts)-1], nil
            }
        }
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

    return "", fmt.Errorf("public key not found in the output")
}

// ReadPassphrase prompts the user to enter a passphrase.
func ReadPassphrase() (string, error) {
    fmt.Print("Enter a unique passphrase for this particular file: ")
    reader := bufio.NewReader(os.Stdin)
    passphrase, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(passphrase), nil
}
