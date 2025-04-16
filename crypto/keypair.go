package crypto

type Keypair struct {
	PrivateKey PrivateKey
	PublicKey  PublicKey
}

func GenerateKeypair() (PrivateKey, PublicKey) {
	privKey, err := GeneratePrivateKey()
	if err != nil {
		panic(err)
	}
	pubKey := privKey.PublicKey()
	return privKey, pubKey
}
