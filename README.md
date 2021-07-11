# goransom
Simple ransomware written in Golang.

I made this program for educational purposes; there's no harm intended. Ransomware is more and more in the media and that triggered my curiosity.

### How goransom works
It iterates through the Windows filesystem, looking for files. Goransom opens the file, encrypts its content with AES-GCM 256, and continues to the next file. In addition, a POST request to the decryption API will be made to ensure that the victim can decrypt his machine after the victim paid the right amount.

The decryption API generates machine-specific executables for victims. This executable is responsible for decrypting the victim's filesystem.

So, the flow is like this:
1. Encrypt the AES key with RSA
2. Send the encrypted AES key to the API
3. Store the encrypted AES key to the filesystem
4. Iterate through the filesystem
5. Encrypt the files

### Technical implementation
TODO