package sss

import (
	"math/big"
	"testing"

	"github.com/jnsoft/jngo/misc"
	. "github.com/jnsoft/jngo/testhelper"
)

func Test_SSS(t *testing.T) {
	t.Run("test sss", func(t *testing.T) {

		secret := big.NewInt(int64(12348354))
		securitylevel := 7
		noOfshares := 6
		min_shares := 4

		test2 := GetSecurityLevel(256)
		println(test2)

		shares, err := CreateShares(secret, noOfshares, min_shares, securitylevel)

		AssertNil(t, err)

		xs := BigIntegerRange(1, noOfshares+1)
		selected_xs := misc.GetRandomElements[*big.Int](xs, min_shares, false)
		selected_shares := make([]*big.Int, min_shares)
		for i, idx := range selected_xs {
			index := int(idx.Int64()) - 1
			selected_shares[i] = shares[index]
		}
		/*
			println("\nshares:")
			for _, share := range shares {
				println(share.String())
			}
			println("\n xs:")
			for _, t := range selected_xs {
				println(t.String())
			}
			println("\n selected shares:")
			for _, t := range selected_shares {
				println(t.String())
			}
		*/

		res_secret := RecoverSecret(selected_shares, selected_xs, securitylevel)

		AssertEqual(t, res_secret.String(), secret.String())

	})

	t.Run("test sss with 256-bit key", func(t *testing.T) {

		noOfshares := 6
		min_shares := 3

		key := misc.GetRandomBytes(256)

		shares, err := CreateSecretsFromKey(key, noOfshares, min_shares)
		AssertNil(t, err)

		selected_shares := misc.GetRandomElements[string](shares, min_shares, false)

		recovered_key, err := GetKeyFromSecrets(selected_shares, GetSecurityLevel(len(key)))
		AssertNil(t, err)

		CollectionAssertEqual(t, recovered_key, key)

	})
}
