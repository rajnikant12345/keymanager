package aes

import (
	"errors"
	"crypto/rand"
	"io"
)

type Aes struct {
	name string
	Size int
	RawBytes []byte
}





func CreateAESKey (name string , size int) (*Aes , error) {

	var message string



	aes := new(Aes)

	aes.RawBytes = make([]byte,size)

	_,err := io.ReadFull(rand.Reader,aes.RawBytes[:])

	if err != nil {
		return nil, errors.New(message)
	}

	aes.name = name
	aes.Size = size

	return aes,nil

}



