package aes

import (
	"crypto/rand"
	"io"
)

type Aes struct {
	name string
	Size int
	RawBytes []byte
}





func CreateAESKey (name string , size int) (*Aes , error) {

	aes := new(Aes)

	aes.RawBytes = make([]byte,size)

	_,err := io.ReadFull(rand.Reader,aes.RawBytes[:])

	if err != nil {
		return nil, err
	}

	aes.name = name
	aes.Size = size

	return aes,nil

}



