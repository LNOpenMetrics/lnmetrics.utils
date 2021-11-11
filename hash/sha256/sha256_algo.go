package sha256

import (
	go_sha256 "crypto/sha256"

	"github.com/LNOpenMetrics/lnmetrics.utils/hash"
)

//Calculate the sha256 of the content, and return back the
// string in bach54
func SHA256(content *string) string {
	strBytes := hash.ToBytes(content)
	hashStr := go_sha256.Sum256(strBytes)
	return hash.ToString(hashStr[:])
}
