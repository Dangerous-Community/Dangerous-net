# Dangerous-net User interface upgrade Readme

## New branch for the development and exploaration of a new UI. 
New functions from other branches should be adapted into it ASAP. Once everything is covered we can merge back into main. 
Basically, the program needs a user interface, we could use some other ones out there though i like bubble tea, downside is, it is a different structure, using MODEL VIEW UPDATE etc instead of a procedual program style as is already written here.
Though with the MVU style it will greatly improve user experience and feel more like a program to use and hang around in than another linux CLI tool. 

See below for image on bubble tea UI GIF. Also link below for the GitHub of bubble tea UI, 

[Bubble Tea Git Hub](https://github.com/charmbracelet/bubbletea)


![68747470733a2f2f73747566662e636861726d2e73682f627562626c657465612f627562626c657465612d6578616d706c652e676966](https://github.com/SATUNIX/Dangerous-net/assets/111553838/a0e5a5b8-2409-4de0-b3ad-7d8c030092c3)


**Rebranding project from IPFSS to Dagnerous-net** 
- To avoid confusion with the IPFS protocol and to also referance the great support from the Dangerous Things community.
- Dangerous-net is not affiliated with, a part of, or owned by Dangerous Things, it is a community project. 

A unique IPFS frontend that you can use to push your files through. Encrypt all traffic with a Asymmetric RSA keypair and GPG. Ensure that you can upload private data to IPFS, and only you can receive and view on another machine. The HTTPS of IPFS.

**Recommended to use keycard, Apex, or Keycard for Multi Factor Authentication (MFA), encrypting and decrypting data. Using IPFS in general.**


Dangerous Net is an application and library set that integrates the InterPlanetary File System (IPFS) with robust RSA encryption in Go, providing a secure method to store and retrieve files. It encrypts files before uploading to IPFS and decrypts them using a corresponding key pair. The project will also be expanding into a self-sufficient decentralised network for messaging, file transfer, cloud like storage, and incentivisation to contribute to the network.  

## For developers / Bounties 

**In order to accelerate development I am accepting bounties for people who can build the below implementations** 
1. **IPFS Kubo Implementation.** > ~~0ETH~~
  - ✅ Complete

2. **Keycard Implementation.** > ~~0ETH~~
  - ✅ Complete
  
3. **Ways to contribute:** 
   - 🟡 Anything in issues is fair game to fix or submit an enhancement for. 

4. **IPFS Cluster Implementation** **0.05 ETH**
   - ✖️ TO DO!! Implement IPFS clustering so that all users of this application can opt in for the Dangerous Net, the IPFS cluster keeping your encrypted files available and ready to use anytime anywhere.
  
5. **GPG Applet Integration** **0.025 ETH**
   - ✖️ Build support for the GPG / PGP applets used in keycard and flexSecure implants.
   - ✖️ Allow for encryption on chip outside of the application and machine itself, use the in-vivo crypto chips. 

## From Dangerous Things with love. 

The Apex Flex and FlexSecure allow us mortal humans to perform cryptographic functions in vivo (under the skin) this fact paired with the above cryptographic MFA, provides the user a secure way to keep their data safe. 

![image](https://github.com/SATUNIX/IPFSS_IPFS-Secure/assets/111553838/c28a0a23-1c19-4e04-b621-ef7b76d92f77)

You may be asking **"but satunix why is this so special?" "These implants can do PGP and OTP!!!"** Well, they sure can, but thats it, good luck loading several applets onto your keycard for each purpose, then trying to navigate and use all of the different block positions keys, algorithims.... and whatever tf. Me personally, I kave a FlexSecure loaded with Keycard. Thus, Keycard must be used for this process. This allows even the noobiest of users ease of control and access. 
*A load and swipe process.*

>"The ability to carry your OTP authenticator, PGP, and other cryptographic keys, and perform cryptographic functions all in vivo (generate OTP codes, encrypt & decrypt data, etc.) without ever revealing private keys to the NFC interface you are interacting with is a huge step forward for personal digital identity and data security."
   
## Key Features

- **Asymmetric Encryption**: Utilize RSA encryption to secure your files. Files are encrypted with a public key and can only be decrypted with the corresponding passphrase protected private key.
- **Symmetric Encryption**: Utilize a Apex or FlexSecure implant with keycard to use Multi Factor Symmetric encryption on your files, supply a passphrase, scan your card, files secured. 
- **Decentralized Storage**: Leverage IPFS for secure, encrypted, decentralized, and immutable file storage.
- **Go Implementation**: Built with Go, taking advantage of its powerful concurrency features and efficient data handling.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go programming language
- IPFS daemon running locally or accessible remotely
- Relevant Go libraries for IPFS and encryption

## Installation


## Contributing

Contributions to SecureIPFS / Dangerous Net are welcome! Please read our upcoming contributing guidelines for details on how to submit pull requests.
For now just create a fork and then create a pull request with your changes.


## License

This project is licensed under the MIT, Apache2 and GPL3 License - see the LICENSE file for details.
Much credit to the IPFS, Keycard, and AES / GPG developers for inspiration and library support for this project. See acknowledgements bellow. 

## Acknowledgments

- IPFS Team for the IPFS protocol
- Go community for the extensive libraries and support
- All contributors who participate in this project

## Roadmap Plan of Development 

**2023**  
- Completion of initial encryption and sys arg functions / data flow to ipfs get and ipfs add.
- New UI, possible bubbletea integration 

**Early 2024** 
- Completion / Signifficant progress towards the Dangerous Net (Building scalable decentralised infrastructure for cloud storage) 
- Usage of JAVA Keycard or other implants for storing the CID log  

**Late 2024**
- Reach of critical mass of the Dangerous Net by application users to support itself and be self sufficient.
- Decentralised network needs to reach a critical *undertermined* size to support itself without daemons running in cloud or in my own servers while also maintaining full file availability.

**Late 2024 to Early 2025**
- Android support

## Donations 

**This project is going to undergo a major overall in networking capability**

[Ethereum Address](https://app.ens.domains/satunix.eth)

Donations are not needed, though are very welcome to support this project and speed up the expansion of the network, all funds sent will directly go towards hardware for this project to support the Dangerous Network. 

![image](https://github.com/SATUNIX/Dangerous-net/assets/111553838/29df7326-7954-4a65-84f0-d191c8ac05e3)


