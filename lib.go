package HashStrings

import (
	"crypto/sha256"
	"fmt"
)

// Single item interface
func Sha256(s string) (rv string) {
	rv = HashStrings(s)
	return
}

// hash a set of strings and return in hex-strings form
func HashStrings(a ...string) string {
	h := sha256.New()
	for _, z := range a {
		h.Write([]byte(z))
	}
	return fmt.Sprintf("%x", (h.Sum(nil)))
}

func HashByte(a []byte) (rv []byte) {
	rv = HashBytes(a)
	return
}

// hash a set of []byte and return in byte form
func HashBytes(a ...[]byte) []byte {
	h := sha256.New()
	for _, z := range a {
		h.Write(z)
	}
	return h.Sum(nil)
}
