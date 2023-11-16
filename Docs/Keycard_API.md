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


## Further Reading for Developers 

Most promising and recomended: **The Official Go API**

- [Keycard for Go applications](https://github.com/status-im/keycard-go/)

Might be the easiest to implement, just have to make sure your can encrypt and decrypt files using the keycard CLI. 

- [Keycard for CLI (can pass sysargs with Go to keycard CLI)](https://github.com/status-im/keycard-cli)

Could use APDU, whatever tf that means. 

- [ Keycard APDU API ](https://keycard.tech/docs/apdu/)https://keycard.tech/docs/apdu/

Or we could use the WEB3 Application and push it through a localhost web UI (One of my favs for the future)

- [Localhost web3 application](https://keycard.tech/docs/web3.html)https://keycard.tech/docs/web3.html


