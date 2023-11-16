# IPFSS_IPFS-Secure
A unique IPFS frontend that you can use to push your files through. Encrypt all traffic with a Asymmetric RSA keypair. Ensure that you can upload private data to IPFS, and only you can receive and view on another machine. The HTTPS of IPFS.


# SecureIPFS

SecureIPFS is an application and library set that integrates the InterPlanetary File System (IPFS) with robust RSA encryption in Go, providing a secure method to store and retrieve files. It encrypts files before uploading to IPFS and decrypts them using a corresponding key pair.

## Key Features

- **Asymmetric Encryption**: Utilize RSA encryption to secure your files. Files are encrypted with a public key and can only be decrypted with the corresponding passphrase protected private key.
- **Decentralized Storage**: Leverage IPFS for secure, decentralized file storage.
- **Go Implementation**: Built with Go, taking advantage of its powerful concurrency features and efficient data handling.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go programming language
- IPFS daemon running locally or accessible remotely
- Relevant Go libraries for IPFS and encryption

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/SecureIPFS.git
   cd SecureIPFS
   ```

2. **Build the Application**:
   ```bash
   go build
   ```
3. Use this instead for IPFS add and get for tunneled transfer.


## Usage

- **Encrypting and Uploading a File**:
  Encrypt your files with a public key and upload them to IPFS, receiving an IPFS hash in return.
  ```bash
  ./SecureIPFS upload --file path/to/file --pubkey path/to/pubkey
  ```

- **Downloading and Decrypting a File**:
  Download and decrypt files from IPFS using your private key.
  ```bash
  ./SecureIPFS download --hash "ipfs_hash" --privkey path/to/privkey --passphrase "your_passphrase"
  ```

## Contributing

Contributions to SecureIPFS are welcome! Please read our upcoming contributing guidelines for details on how to submit pull requests.
For now just create a fork and then create a pull request with your changes.
## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- IPFS Team for the IPFS protocol
- Go community for the extensive libraries and support
- All contributors who participate in this project

## Roadmap Plan of Development 

- Completion of initial encryption and sys arg functions / data flow to ipfs get and ipfs add.
- Usage of JAVA Keycard for signing / backing up private public keys.  

## Function Layout
Key Management Functions

generate_keypair(passphrase): Generates a new RSA key pair (public and private keys) secured by the given passphrase. Stores the key pair securely.
load_private_key(passphrase): Loads the private key from storage, using the passphrase to decrypt it.
get_public_key(): Retrieves the public key from the stored key pair.
File Encryption and Uploading

encrypt_and_sign_file(file_path, private_key): Encrypts the file located at file_path using the private key. Signs the encrypted data to ensure authenticity. Returns the encrypted data.
upload_to_ipfs(encrypted_data): Uploads the encrypted data to IPFS and returns the IPFS hash (CID - Content Identifier) for the uploaded file.
File Downloading and Decryption

download_from_ipfs(ipfs_hash): Downloads the file from IPFS using the given IPFS hash. Returns the downloaded encrypted data.
decrypt_file(encrypted_data, private_key): Decrypts the encrypted data using the private key. Returns the decrypted file content.
Utility Functions

save_file(file_content, file_path): Saves the decrypted file content to the specified path.
read_file(file_path): Reads the content of the file at the specified path.
User Interface / Interaction Functions

These functions would handle user inputs and actions from the GUI, triggering the appropriate backend functions based on user interaction (e.g., button clicks, form submissions).
Example Workflow
Key Generation:

The user generates a new RSA key pair upon first use or loads an existing one.
Upload Workflow:

User selects a file to upload.
The file is encrypted and signed using the user's private key.
The encrypted file is uploaded to IPFS, and the IPFS hash is returned.
Download Workflow:

User inputs an IPFS hash to download the corresponding file.
The encrypted file is downloaded from IPFS.
The user enters the passphrase to unlock the private key.
The downloaded file is decrypted using the private key and saved locally.
Security and User Experience Considerations
Key Security: Store the key pair securely, ideally using an established library for key management.
Passphrase Handling: Ensure the passphrase is not stored insecurely or transmitted in plain text.
Error Handling: Implement robust error handling, especially for cryptographic operations and IPFS interactions.
User Feedback: Provide clear feedback to the user, especially in cases of errors or successful operations.
Conclusion
This function layout provides a structured approach to building your application, focusing on key areas like key management, file encryption/decryption, and interaction with IPFS. It is essential to utilize established cryptographic libraries and follow best practices in security and user interface design to ensure the application is both secure and user-friendly.
