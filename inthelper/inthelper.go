package inthelper

import (
	"math"
	"math/big"

	misc "github.com/jnsoft/jngo"
)

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
	zero := big.NewInt(0)
	for c.Cmp(zero) > 0 {
		a.Mul(&a, big.NewInt(10))
		c.Div(&c, big.NewInt(10))
	}
	return big.NewInt(0).Add(&a, &b)
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
func PseudoPrime(n big.Int) bool {
	return ModularExponentiation(n).Text(10) == "1"
}

func ModularExponentiation(n big.Int) *big.Int {
	ONE := big.NewInt(1)
	base := big.NewInt(2)
	i := big.NewInt(1).Sub(&n, ONE)
	res := n.Exp(base, i, &n)
	return res
}

// FACTORS AND DIVISORS
