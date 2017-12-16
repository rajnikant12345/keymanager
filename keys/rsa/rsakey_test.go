package rsa

import "testing"

func TestCreateRSAKey1024(t *testing.T) {
	a,b := CreateRSAKey("rajni",1024)

	if b != nil {
		t.Fail()
	}

	if a.name != "rajni" {
		t.Fail()
	}


	t.Log(string(a.public))

	t.Log(string(a.private))

}


func TestCreateRSAKey2048(t *testing.T) {
	a,b := CreateRSAKey("rajni",2048)

	if b != nil {
		t.Fail()
	}

	if a.name != "rajni" {
		t.Fail()
	}


	t.Log(string(a.public))

	t.Log(string(a.private))

}


func TestCreateRSAKey4096(t *testing.T) {
	a,b := CreateRSAKey("rajni",4096)

	if b != nil {
		t.Fail()
	}

	if a.name != "rajni" {
		t.Fail()
	}


	t.Log(string(a.public))

	t.Log(string(a.private))

}