package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"os"
	"strings"
)
/*
 *
 *
 *
 * SECURITY NOTE *
 * FIND SECURE HANDLING METHOD FOR CLUSTER CONFIG AND CLUSTER SECRET
 * IT SHOULD BE SECURELY PACKAGED IN THE APPLICATION
 * DEFINE THE CORRECT CLUSTER USER NODE PERMISSIONS
 *
 *
 *
 *
 */


// AddFileToCluster uploads a file to the IPFS Cluster.
// It assumes that the IPFS Cluster peer is already running and configured on the machine.
func AddFileToCluster(filePath string) (string, error) {
	cmd := exec.Command("ipfs-cluster-ctl", "add", filePath)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error adding file to IPFS Cluster: %w", err)
	}

	output := out.String()
	return extractCIDFromClusterOutput(output)
}

// extractCIDFromClusterOutput extracts the CID from the output of the ipfs-cluster-ctl add command.
func extractCIDFromClusterOutput(output string) (string, error) {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "added") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil // Assuming the CID is the second part
			}
		}
	}
	return "", fmt.Errorf("CID not found in IPFS Cluster add output")
}

// InitializeClusterPeer sets up the necessary configuration for a user to join an IPFS Cluster.
// This function assumes the IPFS and IPFS Cluster binaries are installed.
func InitializeClusterPeer(clusterSecret string) error {
    // Initialize IPFS if not already initialized
    if err := initIPFS(); err != nil {
        return fmt.Errorf("error initializing IPFS: %w", err)
    }

    // Set up IPFS Cluster configuration
    return setupClusterConfig(clusterSecret)
}

// initIPFS initializes the local IPFS node.
func initIPFS() error {
    if _, err := os.Stat("~/.ipfs"); os.IsNotExist(err) {
        cmd := exec.Command("ipfs", "init")
        if err := cmd.Run(); err != nil {
            return fmt.Errorf("failed to initialize IPFS: %w", err)
        }
    }
    return nil
}


// setupClusterConfig sets up the IPFS Cluster configuration.
func setupClusterConfig(clusterSecret string) error {
    // Create the IPFS Cluster configuration directory if it doesn't exist
    clusterConfigPath := "~/.ipfs-cluster"
    if _, err := os.Stat(clusterConfigPath); os.IsNotExist(err) {
        os.Mkdir(clusterConfigPath, 0700)
    }

    // Write the cluster secret to the appropriate configuration file
    // This is a simplified example; you may need a more sophisticated configuration setup
    secretConfigPath := fmt.Sprintf("%s/service.json", clusterConfigPath)
    secretConfigContent := fmt.Sprintf(`{"cluster_secret": "%s"}`, clusterSecret)
    if err := os.WriteFile(secretConfigPath, []byte(secretConfigContent), 0600); err != nil {
        return fmt.Errorf("failed to write cluster secret config: %w", err)
    }

    return nil
}
