package rsa

import "crypto/rsa"
import (
	"crypto/rand"
	"encoding/pem"
	"crypto/x509"
	"encoding/asn1"
)

type Rsa struct {
	name string
	Size int
	public []byte
	private []byte
}

func savePEMKey(rsa *Rsa, key *rsa.PrivateKey) {


	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	rsa.private = pem.EncodeToMemory(privateKey)

}


func savePublicPEMKey(rsa *Rsa, pubkey rsa.PublicKey) {
	asn1Bytes, _ := asn1.Marshal(pubkey)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}


	rsa.public = pem.EncodeToMemory( pemkey)
}




func CreateRSAKey (name string , size int) (*Rsa , error) {


	key, err := rsa.GenerateKey(rand.Reader, size)

	if err != nil {
		return nil,err
	}

	publicKey := key.PublicKey

	rsap := new(Rsa)

	savePEMKey(rsap,key)

	savePublicPEMKey(rsap, publicKey)

	rsap.name = name

	return rsap , nil

}
