package aes

import "testing"

func TestCreateAESKey(t *testing.T) {
	a,b := CreateAESKey("rajni",16)

	if b != nil {
		t.Fail()
	}

	if a.name != "rajni" {
		t.Fail()
	}

	if len(a.RawBytes) != 16 {
		t.Fail()
	}

	t.Log(a.RawBytes)

}
