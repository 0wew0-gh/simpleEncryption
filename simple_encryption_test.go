package simpleEncryption

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

var (
	loopCount int    = 10000
	contrast  string = "{\"extraItem\":2,\"key\":[\"jb10=m/zkvpds=1/\",\"/*-+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz\"]}"
)

func TestEncryption(t *testing.T) {
	var (
		se      SimpleEncryption = SimpleEncryption{}
		err     error            = nil
		extra   string           = "1"
		str     string           = "1z3a+bc f"
		enStr   string           = ""
		reStr   string           = ""
		reExtra string           = ""
		j       int              = 0
	)
	println(">>>>> TestEncryption New <<<<<")
	se, err = New(2, "jb10=m/zkvpds=1/", "/*-+0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	if err != nil {
		t.Error("encrypt error:", err)
		return
	}
	println("[extra]:     ", extra, "[str]:     ", str)
	enStr = se.encrypt(str, extra)
	println("[encrypt]:                 ", enStr)
	reStr, reExtra, err = se.decrypt(enStr)
	if err != nil {
		t.Error("decrypt error:", err)
		return
	}
	println("[extra]:     ", reExtra, "[decrypt]: ", reStr)

	println(">>>>> TestEncryption New Json <<<<<")
	se, err = NewJson(contrast)
	if err != nil {
		t.Error("encrypt error:", err)
		return
	}
	println("[extra]:     ", extra, "[str]:     ", str)
	for i := 0; i < loopCount; i++ {
		enStr = se.encrypt(str, extra)
		reStr, reExtra, err = se.decrypt(enStr)
		if err != nil {
			t.Error("decrypt error:", err)
			continue
		}
		j++
	}
	println("[loop count]:", j)
	println("[encrypt]:                 ", enStr)
	println("[extra]:     ", reExtra, "[decrypt]: ", reStr)
}

func TestRAS(t *testing.T) {
	var (
		str   string = "1z3abc f"
		enStr string = ""
		reStr string = ""
		j     int    = 0
	)
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		t.Error(err)
		return
	}
	publicKey := privateKey.PublicKey

	for i := 0; i < loopCount; i++ {
		encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, []byte(str))
		if err != nil {
			t.Error(err)
			return
		}
		enStr = string(encryptedBytes)
		decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
		if err != nil {
			t.Error(err)
			return
		}
		reStr = string(decryptedBytes)
		j++
	}

	println("[loop count]:", j)
	println("[encrypt]:   ", enStr)
	println("[decrypt]:   ", reStr)
}
