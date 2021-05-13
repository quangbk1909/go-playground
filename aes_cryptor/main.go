package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	//inp := `sV2wgLZanexf0OQpaoTQexNPwNNPoSb/awwK/5bE90tPrsyx3d2EvKppX9ddzx6/En+5W16VMd/N0DtJmQNgh+ELS6hAcapI1A5pQkVX/fWlrJxS50AdwONT2u3taRfSPrQ=`
	//passPhraseKYCProd := `GJ182S22CPemlK7gCxNJ`
	passPhraseKYCNonProd := `zxcjvkmn14r97sdfnj1wnfasjilvgaioslvhkqw.f`
	input := "CQl9zNMl2wMmu7asTeJhAv0/5wuFyubmdYH2vLkDiZengvES5dUH67X1Yg0O4ThVYm2wI5ltTHSrzt4gKRmoRLXZaWwjmD58YqDK/kS1go/XTQ/zGsMTuQ6TjcAnMAkpoYA5I5se7nW8or3I3swDJOdoXbSVFiZE+I/dyI7xoiGdI+X/d6O7tlfMq4tdIXgTlVs/5Nv2LFn60obKCEOAL1VYc+7yipxguJShF/z4QBg3QU0pMmJ6O7QoelaawHAtDRMQVqlmQZKtxvpqSysEz5Xhj4Hfvcy8Ct8x9X0Ck00wgaIdDjgzbdzUHA0KmImb4SEuTsJMhrBGxhYoTNtqnI/5npVmS/4VDCF2+R4weCEcAsklYvGbqw7mHZ5V+RmLSL8rJzHVNjjpZsXM7PvnsO9a4fvCNswrfXzkg2grc3W2CRQQOy5kCInxZSXiWG6jAA8aRWqFac3mWpuiNWur0Sl+xvL5OQ5qHQAF6PakwxVcHkikUammU/YhWpUQdEfA5bpaekalKD93IxOe9OsSwxuZk5ujRtyZfZCfnLNiEtOkz4WlcJ7DqErdxwP9sg0hND3zNIxZHpzNZdkPsUE7Bq9mLDOdPTNAm/fjSq7FnKzA6kf+qTdYXFna/nM0hQ5ShwxnnuAxbekzw5RE6koix862OyysW6TeFw6LPxBxOZiVuHCEWxddIumX4kjpeDwszVbNeFGdxeMesJgFmv6oeQnK0kkM65gnT48UNN+c64sN50u6SM2Le+z7EPGAAtD4L+Gm2xeuGEDWby4iyteAF0nuKckvyY/hxk94DC9w7bd5hXQi7WB6yGmFTXXTMbdq2JnyGLOOKVTqjiZPBS3HZeGt6ka4gJ8YmAP3SnZ83MWxn1NiX0xweX+4MAUZlyQU7+vYSUYO9nEh292O1I+09qAm8am2v0Q/ObEjjuHUU7yVwomXH6m3RZB2jzryHLrO77/bJQxH1F1TRsl/z5LfOJRLZJp+w3B3wFh/W+d2KnficYK90i1c0jTtFEw4nbJR3i6p58TqPvw81YHmny90HVN4v1Ke"
	intutByte := base64Decode(input)


	//encryptData, err := Encrypt(input, pass)
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Println(string(encryptData))

	//inp := encryptData
	res, err := Decrypt(intutByte, passPhraseKYCNonProd)
	if err != nil {
		panic(err)
	}

	//event := KafkaEventMessage{
	//	//ID:                         "123312",
	//	//Event:                      "fs-digital-onboarding",
	//	//ServiceCode:                "kyc",
	//	//Timestamp:                  time.Now().UnixNano() / 1000000,
	//	//UserID:                     "1204202101",
	//	//PayloadEncryptionAlgorithm: "AES",
	//	Payload:                    encryptData,
	//}
	//eventJson, err:= json.Marshal(event)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(string(res))

	//err = ioutil.WriteFile("output.json", eventJson,0644)
	//if err != nil {
	//	panic(err)
	//}
}

type KafkaEventMessage struct {
	ID                         string `json:"id,omitempty"`
	RequestID                  string `json:"request_id,omitempty"`
	RefEventID                 string `json:"ref_event_id,omitempty"`
	Event                      string `json:"event,omitempty"`
	ServiceCode                string `json:"service_code,omitempty"`
	Timestamp                  int64  `json:"timestamp,omitempty"`
	UserID                     string `json:"user_id,omitempty"`
	PayloadID                  string `json:"payload_id,omitempty"`
	PayloadEncryptionAlgorithm string `json:"payload_encryption_algorithm,omitempty"`
	Payload                    []byte `json:"payload"`
}

func base64Decode(inp string) []byte {
	ret, err := base64.StdEncoding.DecodeString(inp)
	if err != nil {
		panic(err)
	}
	return ret
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
func Encrypt(data []byte, passphrase string) (encryptedData []byte, err error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText, nil
}
func Decrypt(data []byte, passphrase string) (encryptedData []byte, err error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return
	}
	return plaintext, nil
}
