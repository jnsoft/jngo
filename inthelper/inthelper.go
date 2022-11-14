package inthelper

import "math/big"

func Concat(a uint64, b uint64) uint64 {
	c := b
	for c > 0 {
		a *= 10
		c /= 10
	}
	return a + b
}

func Concat_big(a big.Int, b big.Int) *big.Int {
	c := b
	zero := big.NewInt(0)
	for c > zero {
		a.Mul(&a, big.NewInt(10))
		c.Div(&c, big.NewInt(10))
	}
	return big.NewInt(0).Add(&a, &b)
}
