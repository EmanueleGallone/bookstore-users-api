package crypto_util

import (
	"crypto/sha256"
	"encoding/hex"
)

const salt = "super duper secret to salt every password with it! Maybe unneeded? who knows?!"

func GetSHA256Hash(input string) string {
	input += salt
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:]) //hash[:] converts fixe array to slice
}
