package hash

import (
	"encoding/hex"
)

// Util function to wrap the logic to convet a string into a list
// of bytes
func ToBytes(content *string) []byte {
	return []byte(*content)
}

// Util function to wrap the logic to convert a list of bytes
// into a string
func ToString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}
