package string

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/thanhpk/randstr"
)

func GenerateSerial(prefix string, n int) string {
	return fmt.Sprintf("%s%s", prefix, randstr.Hex(n))
}

func GenerateTransactionID() string {
	length := 32
	bytes := make([]byte, length/2)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
