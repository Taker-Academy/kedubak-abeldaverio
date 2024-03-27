package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash_sring(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return (hex.EncodeToString(h.Sum(nil)))
}
