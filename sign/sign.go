// Package sign give the API to implement a custom signer
// that is able to sign and verify a sequence of bytes or
// a string.
package sign

// Signer is a common interface that give the possibility to implement
// a signer and verify algorithm in pure go
type Signer interface {
	// SignMsg sign a string and return the result or an error
	SignMsg(msg *string) (*string, error)

	// VerifyMsg a signature with optional public key specified as parameter
	// or use some key stored inside the struct.
	VerifyMsg(signature *string) (bool, error)
}

// Keys tagger interface to store the information about the
// key (if any) that are used to sign a message.
type Keys interface{}
