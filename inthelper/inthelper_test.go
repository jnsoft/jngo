package inthelper

import (
	"math/big"
	"testing"

	. "github.com/jnsoft/jngo/testhelper"
)

func TestSum(t *testing.T) {
	t.Run("sum of ints", func(t *testing.T) {
		ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		tot := Sum(ns)
		AssertEqual(t, tot, 45)
	})
}

func TestPseudoPrime(t *testing.T) {
	t.Run("PseudoPrime", func(t *testing.T) {
		n, _ := new(big.Int).SetString("218882428714186575612", 0)
		n2, _ := new(big.Int).SetString("785969971488174033889231946017", 0)
		AssertFalse(t, PseudoPrime(n))
		AssertTrue(t, PseudoPrime(n2))
	})
}

func TestModExp(t *testing.T) {
	t.Run("modular exponentiation", func(t *testing.T) {
		n := big.NewInt(37)
		one := big.NewInt(1)
		tot := ModularExponentiation(n)
		AssertEqual(t, one.Text(10), tot.Text(10))
	})
}

func TestMillerRabin(t *testing.T) {
	t.Run("Miller Rabin primality test", func(t *testing.T) {
		p, _ := new(big.Int).SetString("785969971488174033889231946017", 0)
		test1 := MillerRabin(p, 300)
		AssertTrue(t, test1)
	})
}
