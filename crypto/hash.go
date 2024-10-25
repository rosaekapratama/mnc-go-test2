package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(input string) (output string) {
	// Create a new SHA-1 hash
	hash := sha1.New()

	// Write data to the hash
	hash.Write([]byte(input))

	// Calculate the hash
	hashBytes := hash.Sum(nil)

	// Convert to hex string
	output = hex.EncodeToString(hashBytes)
	return
}
