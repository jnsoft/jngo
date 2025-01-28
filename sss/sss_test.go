package sss

import (
	"math/big"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func Test_SSS(t *testing.T) {
	t.Run("test sss", func(t *testing.T) {
		
		secret := 12348354
		securitylevel := 13
		noOfshares := 6
		min_shares := 3

		shares, err := CreateShares(secret, noOfshares, min_shares, securitylevel)
		if err != nil {
			println(err.Error())
		} else {
			for _, share := range shares {
				println(share.String())
			}
		}

		xs := []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)}
		secret := RecoverSecret(shares, xs, 2)
		println(secret.String())

		AssertTrue(t, true)

	})
}
