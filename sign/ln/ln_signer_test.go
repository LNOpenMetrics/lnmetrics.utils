package ln

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test the following command in core lightning
// clightning --testnet signmessage "Testing go utils"
// pub key: 03b39d1ddf13ce486de74e9e44e0538f960401a9ec75534ba9cfe4100d65426880
// zbase32 encoding: d7m5xma1tia1hw6mjkeixqoocexiyoyn1jgm53mncszdeut5j9r1k1nq7wsrh3pkk3c3xp7mkuuxmbneitu3us36cqfn9a6sdu3wa45j
// signature: 57b7af128d712e53cb4a9157ba10621f504002924cbde56265ae344e3b4fc925484eed2c4e65aa565997b7ab54e6f58448ac6799db3e638a2fe3d61cf34c6b69
// message: "Testing go utils"
func TestLNSignerVerifyMsgTrue(t *testing.T) {
	pubKey := "03b39d1ddf13ce486de74e9e44e0538f960401a9ec75534ba9cfe4100d65426880"
	zbase32 := "d7m5xma1tia1hw6mjkeixqoocexiyoyn1jgm53mncszdeut5j9r1k1nq7wsrh3pkk3c3xp7mkuuxmbneitu3us36cqfn9a6sdu3wa45j"
	msg := "Testing go utils"
	signer := NewLNSigner()
	verified, err := signer.VerifyMsg(&pubKey, &zbase32, &msg)
	assert.Nil(t, err)
	assert.Equal(t, true, verified, "Return level returned it is diffirent")
}

func TestLNSignerVerifyMsgFalse(t *testing.T) {
	pubKey := "03b39d1ddf13ce486de74e9e44e0538f960401a9ec75534ba9cfe4100d65426880"
	zbase32 := "d7m5xma1tia1hw6mjkeixqoocexiyoyn1jgm53mncszdeut5j9r1k1nq7wsrh3pkk3c3xp7mkuuxmbneitu3us36cqfn9a6sdu3wa45j"
	msg := "Invalid message"
	signer := NewLNSigner()
	verified, err := signer.VerifyMsg(&pubKey, &zbase32, &msg)
	assert.Nil(t, err)
	assert.Equal(t, false, verified, "Return level returned it is diffirent")
}
