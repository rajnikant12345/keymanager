package rsa

import "testing"

func TestCreateRSAKey1024(t *testing.T) {
	a,b := CreateRSAKey("rajni",1024)

	if b != nil {
		t.Fail()
	}

	if a.Name != "rajni" {
		t.Fail()
	}


	t.Log(string(a.Public))

	t.Log(string(a.Private))

}


func TestCreateRSAKey2048(t *testing.T) {
	a,b := CreateRSAKey("rajni",2048)

	if b != nil {
		t.Fail()
	}

	if a.Name != "rajni" {
		t.Fail()
	}


	t.Log(string(a.Public))

	t.Log(string(a.Private))

}


func TestCreateRSAKey3072(t *testing.T) {
	a,b := CreateRSAKey("rajni",3072)

	if b != nil {
		t.Fail()
	}

	if a.Name != "rajni" {
		t.Fail()
	}


	t.Log(string(a.Public))

	t.Log(string(a.Private))

}



func TestCreateRSAKey4096(t *testing.T) {
	a,b := CreateRSAKey("rajni",4096)

	if b != nil {
		t.Fail()
	}

	if a.Name != "rajni" {
		t.Fail()
	}


	t.Log(string(a.Public))

	t.Log(string(a.Private))

}