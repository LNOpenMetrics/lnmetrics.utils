// Package ln is a package that implement a signature proposed in the
// lightning network protocol, but not yet specified.
package ln

import (
	"encoding/base32"
	"encoding/hex"
	"github.com/LNOpenMetrics/lnmetrics.utils/goutils"
	"github.com/LNOpenMetrics/lnmetrics.utils/hash/sha256"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

// LNSigner is a struct that implement th Signer interface
type LNSigner struct {
	encoder         *base32.Encoding
	signedMsgPrefix string
}

func NewLNSigner() *LNSigner {
	return &LNSigner{
		// Copied from https://github.com/ElementsProject/lightning/blob/master/lightningd/signmessage.c#L11
		encoder:         base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769"),
		signedMsgPrefix: "Lightning Signed Message:",
	}
}

// SignMsg sign the message by following the lightning network rules, but it is not implemented
// because required to import the keys to sign the message.
func (self *LNSigner) SignMsg(msg *string) (*string, error) {
	return nil, goutils.NotImplementYet()
}

// VerifyMsg sign the message with the lightning network logics implemented in the most popular implementation
func (self *LNSigner) VerifyMsg(key *string, signature *string, msg *string) (bool, error) {
	u8sing, err := self.encoder.DecodeString(*signature)
	if err != nil {
		return false, nil
	}

	// https://github.com/ElementsProject/lightning/blob/master/lightningd/signmessage.c#L177
	if len(u8sing) != 65 {
		return false, goutils.Errf("zbase is too is wrong size %d, need to be exactly 65", len(u8sing))
	}

	// The signature is over the double-sha256 hash of the message.
	toVerify := append([]byte(self.signedMsgPrefix), []byte(*msg)...)
	digest := sha256.DoubleSHA256FromByte(toVerify)

	// RecoverCompact both recovers the pubkey and validates the signature.
	pubKey, _, err := ecdsa.RecoverCompact(u8sing, digest)
	if err != nil {
		return false, nil
	}
	pubKeyHex := hex.EncodeToString(pubKey.SerializeCompressed())
	return pubKeyHex == *key, goutils.NotImplementYet()
}
