package sss

import (
	"encoding/base64"
	"errors"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/jnsoft/jngo/misc"
	"github.com/jnsoft/jngo/stringhelper"
)

const (
	DELIMETER = "|"
)

// Shamir's Secret Sharing
type SecretSharing struct{}

// Mersenne primes exponents, e.g. 2^127-1 for desired security level of 128.
// Too large and all the ciphertext is large, too small and security is compromised
var SecurityLevels = []int{5, 7, 13, 17, 19, 31, 61, 89, 107, 127,
	521, 607, 1279, 2203, 2281, 3217, 4253, 4423, 9689, 9941,
	11213, 19937, 21701, 23209, 44497, 86243, 110503, 132049, 216091, 756839,
	859433, 1257787, 1398269, 2976221, 3021377, 6972593, 13466917, 20996011, 24036583, 25964951,
	30402457, 32582657, 37156667, 42643801, 43112609}

func GetPrime(securityLevel int) *big.Int {
	return new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(SecurityLevels[securityLevel])), nil), big.NewInt(1))
}

func CreateSecretsFromKey(key []byte, no_of_shares, threshold int) ([]string, error) {
	secret := bytesToBigInt(key)
	shares, err := CreateShares(secret, no_of_shares, threshold, GetSecurityLevel(len(key)))
	if err != nil {
		return nil, errors.New("CreateSecretsFromKey->CreateShares failed")
	}

	b64_shares := BigIntsToBase64Strings(shares)
	x := 0
	b64_shares_with_x := misc.Map(b64_shares, func(s string) string {
		x++
		return strconv.Itoa(x) + DELIMETER + s
	})
	return b64_shares_with_x, nil
}

func GetKeyFromSecrets(secrets []string, secLevel int) ([]byte, error) {
	secret_strings := stringhelper.SplitStrings(secrets, DELIMETER)
	xs := misc.Map(secret_strings, func(slice []string) *big.Int {
		bigInt := new(big.Int)
		bigInt.SetString(slice[0], 10)
		return bigInt
	})
	y_strings := misc.Map(secret_strings, func(slice []string) string {
		return slice[1]
	})
	ys, err := Base64StringsToBigInts(y_strings)
	if err != nil {
		return nil, errors.New("GetKeyFromSecrets->Base64StringsToBigInts failed")
	}
	secret := RecoverSecret(ys, xs, secLevel)
	key := bigIntToBytes(secret)
	return key, err
}

func CreateShares(secret *big.Int, shares, threshold, securityLevel int) ([]*big.Int, error) {
	if secret.Cmp(big.NewInt(0)) <= 0 {
		return nil, errors.New("secret must be greater than 0")
	}

	if shares < threshold {
		return nil, errors.New("number of shares must not be less than threshold")
	}

	prime := GetPrime(securityLevel)

	// Create the polynomial
	polynomial := []*big.Int{secret}
	for i := 0; i < threshold-1; i++ {
		randomCoeff := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(prime, big.NewInt(1)))
		polynomial = append(polynomial, randomCoeff)
	}

	// Create the shares
	points := make([]*big.Int, shares)
	for i := 1; i <= shares; i++ {
		x := big.NewInt(int64(i))
		y := big.NewInt(0)
		for j := 0; j < len(polynomial); j++ {
			term := new(big.Int).Mul(polynomial[j], new(big.Int).Exp(x, big.NewInt(int64(j)), prime))
			y = new(big.Int).Add(y, term)
		}
		y = new(big.Int).Mod(y, prime)
		points[i-1] = y
	}

	return points, nil
}

// RecoverSecret recovers the secret using Lagrange interpolation
func RecoverSecret(shares, xs []*big.Int, securityLevel int) *big.Int {
	prime := GetPrime(securityLevel)
	return lagrangeInterpolate(big.NewInt(0), xs, shares, prime)
}

// LagrangeInterpolate performs Lagrange interpolation
// Find the y-value for the given x, given n (x, y) points;
// k points will define a polynomial of up to k-1 th order.
// xs must be distinct points

func lagrangeInterpolate(x *big.Int, xs, ys []*big.Int, p *big.Int) *big.Int {
	result := big.NewInt(0)
	for i := 0; i < len(xs); i++ {
		term := new(big.Int).Set(ys[i])
		for j := 0; j < len(xs); j++ {
			if i != j {
				num := new(big.Int).Sub(x, xs[j])
				den := new(big.Int).Sub(xs[i], xs[j])
				denInv := new(big.Int).ModInverse(den, p)
				term.Mul(term, num).Mul(term, denInv).Mod(term, p)
			}
		}
		result.Add(result, term).Mod(result, p)
	}
	return result
}

func BigIntegerRange(start, end int) []*big.Int {
	var result []*big.Int
	for i := start; i < end; i++ {
		result = append(result, big.NewInt(int64(i)))
	}
	return result
}

func GetSecurityLevel(secretSizeInBytes int) int {
	for i := 0; i < len(SecurityLevels)-1; i++ {
		if SecurityLevels[i] < secretSizeInBytes*8 {
			continue
		} else {
			return i
		}
	}
	return len(SecurityLevels) - 1
}

func bytesToBigInt(b []byte) *big.Int {
	bigInt := new(big.Int)
	bigInt.SetBytes(b)
	return bigInt
}

func bigIntToBytes(bi *big.Int) []byte {
	return bi.Bytes()
}

func BigIntsToBytesArray(bigInts []*big.Int) [][]byte {
	bytesArray := make([][]byte, len(bigInts))
	for i, bigInt := range bigInts {
		bytesArray[i] = bigInt.Bytes()
	}
	return bytesArray
}

func BytesArrayToBigInts(bytesArray [][]byte) []*big.Int {
	bigInts := make([]*big.Int, len(bytesArray))
	for i, bytes := range bytesArray {
		bigInt := new(big.Int).SetBytes(bytes)
		bigInts[i] = bigInt
	}
	return bigInts
}

func BigIntsToBase64Strings(bigInts []*big.Int) []string {
	base64Strings := make([]string, len(bigInts))
	for i, bigInt := range bigInts {
		bytes := bigInt.Bytes()
		base64Strings[i] = base64.StdEncoding.EncodeToString(bytes)
	}
	return base64Strings
}

func Base64StringsToBigInts(base64Strings []string) ([]*big.Int, error) {
	bigInts := make([]*big.Int, len(base64Strings))
	for i, str := range base64Strings {
		bytes, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			return nil, err
		}
		bigInt := new(big.Int).SetBytes(bytes)
		bigInts[i] = bigInt
	}
	return bigInts, nil
}

func getFirstElement(slice []string) string {
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}
