package cluster_link

import (
	"bytes"
	"fmt"
	"os/exec"
	"os"
	"strings"
	"encoding/json"
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
type ClusterConfig struct {
    Cluster struct {
        Secret string `json:"secret"`
        // ... other fields from service.json ...
    } `json:"cluster"`
}
const ServiceJsonCID = "QmcDBkxjNL14fmzv46A9xpcGgA4JLgA9ujgScKfR3AzxAx"


func AddFileToCluster(filePath string) (string, error) {
	cmd := exec.Command("ipfs-cluster-ctl", "add", filePath)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error adding file to IPFS Cluster: %w", err)
	}

	output := out.String()
	return ExtractCIDFromClusterOutput(output)
}

// extractCIDFromClusterOutput extracts the CID from the output of the ipfs-cluster-ctl add command.
func ExtractCIDFromClusterOutput(output string) (string, error) {
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
/*
 * See the following for remote joining and cluster setup to be built in:
 * https://ipfscluster.io/documentation/collaborative/setup/
 * https://ipfscluster.io/documentation/deployment/
 * https://ipfscluster.io/documentation/deployment/bootstrap/
 *
 *
 */

func InitializeClusterPeer(clusterSecret string) error {
    // Initialize IPFS if not already initialized
    if err := InitIPFS(); err != nil {
        return fmt.Errorf("error initializing IPFS: %w", err)
    }
    return SetupClusterConfig(clusterSecret)
}

func InitIPFS() error {
    if _, err := os.Stat("~/.ipfs"); os.IsNotExist(err) {
        cmd := exec.Command("ipfs", "init")
        var stderr bytes.Buffer
        cmd.Stderr = &stderr

        if err := cmd.Run(); err != nil {
            if strings.Contains(stderr.String(), "someone else has the lock") {
                fmt.Println("IPFS daemon is already running.")
                return nil
            }
            return fmt.Errorf("failed to initialize IPFS: %w", err)
        }
    }
    return nil
}


func RetrieveAndApplyClusterConfig() error {
    // Retrieve service.json from IPFS using its CID
    cmd := exec.Command("ipfs", "get", ServiceJsonCID, "-o", "service.json")
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to retrieve cluster configuration from IPFS: %w", err)
    }

    // Read the configuration file
    configFileData, err := os.ReadFile("service.json")
    if err != nil {
        return fmt.Errorf("failed to read cluster config file: %w", err)
    }

    // Parse the JSON configuration
    var config ClusterConfig
    if err := json.Unmarshal(configFileData, &config); err != nil {
        return fmt.Errorf("failed to parse cluster config: %w", err)
    }

    // Use the extracted secret to set up the cluster configuration
    if err := SetupClusterConfig(config.Cluster.Secret); err != nil {
        return err
    }

    return StartClusterDaemon()
}

func StartClusterDaemon() error {
    cmd := exec.Command("ipfs-cluster-service", "daemon")
    var stderr bytes.Buffer
    cmd.Stderr = &stderr

    if err := cmd.Start(); err != nil {
        if strings.Contains(stderr.String(), "someone else has the lock") {
            fmt.Println("IPFS Cluster daemon is already running.")
            return nil
        }
        return fmt.Errorf("failed to start IPFS Cluster daemon: %w", err)
    }
    return cmd.Wait()
}

// setupClusterConfig sets up the IPFS Cluster configuration.
func SetupClusterConfig(clusterSecret string) error {
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
