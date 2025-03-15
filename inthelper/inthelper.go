package inthelper

import (
	"crypto/rand"
	"math"
	"math/big"

	misc "github.com/jnsoft/jngo/misc"
)

var ZERO *big.Int = big.NewInt(0)
var ONE *big.Int = big.NewInt(1)
var TWO *big.Int = big.NewInt(2)

// HELPERS

func Sum(arr []int) int {
	return misc.Fold(arr, func(a, b int) int {
		return a + b
	}, 0)
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

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
	for c.Cmp(ZERO) > 0 {
		a.Mul(&a, big.NewInt(10))
		c.Div(&c, big.NewInt(10))
	}
	return big.NewInt(0).Add(&a, &b)
}

func Reverse(n int) int {
	rev := 0
	for n > 0 {
		rev = 10*rev + n%10
		n = n / 10
	}
	return rev
}

// Greatest Common Divisor (a.k.a. gcd, gcf, hcf, gcm, sgd (in swedish))
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Least Common Multiple via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func IsPerfectPower(n int, power int) (bool, int) {
    if n < 0 && power%2 == 0 {
        return false, 0 // Negative numbers cannot have real even roots
    }
    root := int(math.Round(math.Pow(float64(n), 1.0/float64(power))))
    // Verify that root^power equals n
    if int(math.Pow(float64(root), float64(power))) == n {
        return true, root
    }
    return false, 0
}

func Permutations(n int) []int {
    str := strconv.Itoa(n)
    digits := []rune(str)

    permutations := []string{}
    permute(digits, 0, &permutations)

    result := []int{}
    for _, perm := range permutations {
        num, _ := strconv.Atoi(perm)
        result = append(result, num)
    }
    return result
}

func permute(digits []rune, start int, permutations *[]string) {
    if start == len(digits)-1 {
        *permutations = append(*permutations, string(digits))
        return
    }

    for i := start; i < len(digits); i++ {
        // Swap current element with the starting element
        digits[start], digits[i] = digits[i], digits[start]

        // Recursively generate permutations for the remaining digits
        permute(digits, start+1, permutations)

        // Swap back to restore the original state
        digits[start], digits[i] = digits[i], digits[start]
    }
}

// PRIMES

func PrimesSieve(exclusive_limit int) []int {
	var sievebound int = exclusive_limit / 2 // last index of sieve
	var sieve = make([]bool, sievebound)
	for i := range sieve {
		sieve[i] = true
	}
	crosslimit := int(math.Floor(math.Sqrt(float64(exclusive_limit))-1)/2 + 1)
	for i := 1; i < crosslimit; i++ {
		if sieve[i] { // 2*i+1 is prime, mark multiples
			for j := 2 * i * (i + 1); j < sievebound; j += 2*i + 1 {
				sieve[j] = false
			}
		}
	}

	var res = make([]int, 0)
	for i, val := range sieve {
		if val {
			res = append(res, 2*i+1)
		}
	}
	res[0] = 2
	return res
}

// Daisy-chain Filter processes
func PrimesSieveConcurent(no_of_primes int) []int {
	var res = make([]int, no_of_primes)
	ch := make(chan int)             // create a new channel
	go generate_prime_candidates(ch) // launch Generate goroutine.
	for i := 0; i < no_of_primes; i++ {
		prime := <-ch
		// println("new prime, adding filter for " + strconv.Itoa(prime))
		res[i] = prime
		if i == no_of_primes-1 {
			break
		} else {
			ch_next := make(chan int)
			go filter_primes(ch, ch_next, prime)
			ch = ch_next // the input of the next filter is the output of this one
		}
	}
	return res
}

func generate_prime_candidates(ch chan<- int) {
	for i := 2; ; i++ {
		// println("sending " + strconv.Itoa(i))
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out', removing those divisible by 'prime'.
func filter_primes(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		// println("recieved " + strconv.Itoa(i) + " to " + strconv.Itoa(prime) + " filter")
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

/*
func IsPrime(n:int):
    if (n <= 1):
        return False
    elif (n < 4):
        return True # 2,3
    elif (n % 2 == 0):
        return False
    elif (n < 9):
        return True # 5,7
    elif (n % 3 == 0):
        return False
    elif(pseudo_prime(n)):
            if (miller_rabin(n, 40)):
                return True
    return False
*/

// pseudo primality testing using Fermat's theorem,
// can report false primes, but never false negatives
func PseudoPrime(n *big.Int) bool {
	return ModularExponentiation(n).Text(10) == "1"
}

func ModularExponentiation(n *big.Int) *big.Int {
	i := new(big.Int).Sub(n, ONE)
	res := new(big.Int).Exp(TWO, i, n)
	return res
}

// MillerRabin checks if n is a probable prime using k iterations of the Miller-Rabin test
func MillerRabin(n *big.Int, k int) bool {

	if n.Cmp(TWO) == 0 {
		return true // 2 is a prime number
	}

	if n.Cmp(ONE) <= 0 || n.Bit(0) == 0 {
		return false // n <= 1 or even
	}

	// write n as 2^s * d + 1 with d odd (by factoring out powers of 2 from n − 1)
	//s := big.NewInt(0)
	s := 0
	d := new(big.Int).Sub(n, ONE)

	for new(big.Int).Mod(d, TWO).Cmp(ZERO) == 0 {
		d.Div(d, TWO)
		s++
	}

	for i := 0; i < k; i++ { // witness loop
		a, _ := rand.Int(rand.Reader, new(big.Int).Sub(n, TWO))
		a.Add(a, ONE) // Ensure a is in the range [1, n-1]
		x := new(big.Int).Exp(a, d, n)
		if x.Cmp(ONE) == 0 || x.Cmp(new(big.Int).Sub(n, ONE)) == 0 { // # n is always a probable prime to base 1 and n − 1
			continue
		}
		for j := 0; j < s-1; j++ { // inner loop
			x.Exp(x, TWO, n)
			if x.Cmp(ONE) == 0 {
				return false
			}
			if x.Cmp(new(big.Int).Sub(n, ONE)) == 0 {
				break
			}
		}
		if x.Cmp(new(big.Int).Sub(n, ONE)) != 0 {
			return false
		}
	}
	return true
}

// FACTORS AND DIVISORS

func GetFactor(n int) (int, int) {
	if MillerRabin(big.NewInt(int64(n)), 13) {
		return n, 1
	}
	if n % 2 == 0 {
		return 2, n / 2
	}
	if n % 3 == 0 {
		return 2, n / 3
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 5; i <= limit; i+=6 {
		if n%i == 0 {
			return i, n / i
		} else if n % (i + 2) == 0 {
                        return i+2, n / (i+2)
			}
			
	}
	return n,-1
}

func Factor(n int) []int {
	factors := []int{}
	var factor int

	for n > 1 {
		factor, n = GetFactor(n)
		factors = append(factors, factor)
	}
	return factors
}
