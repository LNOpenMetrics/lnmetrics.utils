package sha256

import (
	go_sha256 "crypto/sha256"

	"github.com/LNOpenMetrics/lnmetrics.utils/hash"
)

// SHA256 Calculate the sha256 of the content, and return back the
// string in bach54
func SHA256(content *string) string {
	strBytes := hash.ToBytes(content)
	hashStr := go_sha256.Sum256(strBytes)
	return hash.ToString(hashStr[:])
}

// DoubleSHA256 Double hash is the algorithm used in Bitcoin protocol
func DoubleSHA256(content *string) string {
	strBytes := hash.ToBytes(content)
	oneHash := go_sha256.Sum256(strBytes)
	hashStr := go_sha256.Sum256(oneHash[:])
	return hash.ToString(hashStr[:])
}

// DoubleSHA256FromByte Double hash is the algorithm used in Bitcoin protocol
func DoubleSHA256FromByte(content []byte) []byte {
	oneHash := go_sha256.Sum256(content)
	hashStr := go_sha256.Sum256(oneHash[:])
	return hashStr[:]
}
