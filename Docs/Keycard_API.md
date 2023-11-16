# Keycard implementation 

## Notes

Understanding Keycard Capabilities:
1. Key Storage:
- Keycard is designed to securely store private keys, typically used for digital signatures and authenticating transactions, especially in blockchain contexts.

3. Signing Operations:
- Keycard can sign data using the stored private keys. This is commonly used for signing transactions in cryptocurrency applications. Instead we are using a file based application where instead of blocks of cryptocurrency being signed it is blocks of IPFS storage. 
- The idea would be akin to keeping private public keys for Ethereum / ERC-20. Something which is well proven and battle tested for keycard through Status.
- Instead we just decrypt the data using the private keycard key. ***Though this will be different to the core of the IPFSS program as the private key is not revealed to the interface itself.*** 

5. Encryption/Decryption:
- Standard file encryption/decryption is not a primary function of Keycard. It's more oriented towards signing and transaction authentication.
- Though this signing and transaction authentication can be expanded upon. Since cryptocurrency transactions are just private key signings this shouldnt be too hard to implement. 


## Keycard Implementation for Encryption and Decryption

### Utilize Keycard for Signing:

- The Keycard will be used for signing data blocks, similar to how it signs cryptocurrency transactions.
- For file-based applications, each block of data (or the entire file) can be signed using the Keycard's private key.

### Generating and Storing Keys:

- Use commands like `GENERATE KEY` or `LOAD KEY` to create or load RSA keys onto the Keycard.
- The generated keys can be used for signing operations. Ensure these keys are securely stored and managed on the Keycard.

### File Encryption and Decryption:

#### Encryption:
- Encrypt files using standard encryption tools, like OpenSSL or GPG, with a public key. This public key can be the one associated with the Keycard's private key, or another keypair depending on your security design.

#### Decryption:
- Decryption is typically performed using the corresponding private key. However, as Keycard does not directly expose private keys for security reasons, it cannot directly decrypt data in the conventional sense. Instead, consider using Keycard for validating the signature of the encrypted data to ensure its integrity and authenticity.
- Alternative Approach: If direct decryption with Keycard's private key is necessary, you may need to explore if `EXPORT KEY` can be used securely, though this might pose a security risk as exporting private keys is generally discouraged.

### Signing Encrypted Files:

- After encrypting a file, use the `SIGN` command to sign the encrypted file with the Keycard's private key.
- This signature can be used to verify the file's integrity and authenticity upon decryption.

### Verifying Signatures:

- When a file is decrypted using a standard decryption tool, use Keycard’s public key to verify the signature. This ensures that the file was indeed encrypted and signed by the owner of the Keycard.

### Workflow Integration:

- Implement a workflow in your application where every file to be uploaded to IPFS is first encrypted, then signed using the Keycard.
- For downloading, after retrieving the file from IPFS, the application should first verify the signature using Keycard's public key, then decrypt it using the corresponding private key (not stored on Keycard).

### Security Protocols:

- Ensure all interactions with Keycard, like PIN verification (`VERIFY PIN`), opening secure channels (`OPEN SECURE CHANNEL`), and mutual authentication (`MUTUALLY AUTHENTICATE`), are handled securely in your application.

### Conclusion

While Keycard provides robust capabilities for signing and key management, its direct use in standard file encryption/decryption workflows is limited due to its design as a secure element for transaction signing and authentication. Your implementation can leverage Keycard for signing encrypted files, adding a layer of security by ensuring data integrity and origin verification. However, for conventional encryption and decryption of files, reliance on external tools and methodologies that complement Keycard's capabilities would be necessary.




## Further Reading for Developers 

Most promising and recomended: **The Official Go API**

- [Keycard for Go applications](https://github.com/status-im/keycard-go/)

Might be the easiest to implement, just have to make sure your can encrypt and decrypt files using the keycard CLI. 

- [Keycard for CLI (can pass sysargs with Go to keycard CLI)](https://github.com/status-im/keycard-cli)

Could use APDU, whatever tf that means. 

- [ Keycard APDU API ](https://keycard.tech/docs/apdu/)https://keycard.tech/docs/apdu/

Or we could use the WEB3 Application and push it through a localhost web UI (One of my favs for the future)

- [Localhost web3 application](https://keycard.tech/docs/web3.html)https://keycard.tech/docs/web3.html


