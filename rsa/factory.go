package rsa

func CreateRSAEncrypterFromKey(publicKey []byte) (Encrypter, error) {
	block, _ := pem.Decode(publicKey)

	parsedPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return Encrypter{PublicKey: parsedPublicKey}, nil
}
