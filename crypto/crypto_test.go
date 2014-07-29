package crypto_test

import (
	"encoding/json"
	"fmt"

	"crypto/rand"
	"encoding/base64"

	// "crypto/sha256"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func CryptoNonce() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err.Error())
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func EncodeObject(payload interface{}) (string, error) {
	extra := map[string]string{"nonce": CryptoNonce()}
	load := []interface{}{payload, extra}

	bytes, err := json.Marshal(load)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// func SignObject(object []byte, priv *PrivateKey) (s []byte, err error) {
// 	func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) (s []byte, err error)
// }

var _ = Describe("Crypto Stuff", func() {
	baselineProperties := map[string]string{"user_id": "42", "roles": "fat,dumb,happy"}

	It("Marshals as a string", func() {
		str, err := EncodeObject(baselineProperties)
		Expect(err).To(BeNil())
		fmt.Println(str)
	})

})
