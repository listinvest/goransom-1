# goransom
Simple ransomware written in Golang.

I made this program for educational purposes; there's no harm intended. Ransomware is more and more in the media and that triggered my curiosity.

### How goransom works
It iterates through the Windows filesystem, looking for files. Goransom opens the file, encrypts its contents with AES-GCM 256, and continues to the next file. A POST request to the 'decryption' service will be made, containing the AES key encrypted with RSA.

So, the flow is like this:
1. Encrypt the AES key with RSA
2. Send the encrypted AES key to the API
3. Store the encrypted AES key to the filesystem
4. Iterate through the filesystem
5. Encrypt the files

### Technical implementation
TODO