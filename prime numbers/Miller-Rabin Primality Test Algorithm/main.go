package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Miller-Rabin Primality Test Algorithm
// The Miller-Rabin primality test is a probabilistic algorithm used to determine whether a given number is likely to be prime.
// It is faster than deterministic primality tests like the extended Euclidean algorithm,
// but there is a small chance it could declare a composite number as prime.
// The Miller-Rabin test is probabilistic, not deterministic.
// There is a small chance (less than 4^-k) that it will misclassify a composite number as prime.
// It doesn't work for all composite numbers. Some composite numbers (called Carmichael numbers) will pass the test regardless of the base a.
// Wikipedia article on Miller-Rabin primality test: https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test
// Interactive visualization of the algorithm: https://www.primegrid.com/
//
// n: The number to test for primality.
// k: The number of iterations (determines the accuracy).
// More iterations (higher k) increase the accuracy of the test but also take longer.
// The algorithm works by checking if a randomly chosen base a exhibits certain properties when raised to different powers modulo n.
// If these properties hold, it suggests that n is likely prime.
func isPrimeMillerRabin(n *big.Int, k int) bool {
	zero := big.NewInt(0)
	two := big.NewInt(2)
	three := big.NewInt(3)

	// If n <= 1 or n is divisible by 2 or 3
	if n.Cmp(two) <= 0 || new(big.Int).Mod(n, two).Cmp(zero) == 0 || new(big.Int).Mod(n, three).Cmp(zero) == 0 {
		return n.Cmp(two) == 0 || n.Cmp(three) == 0
	}

	// Express n - 1 as 2^s * d:
	// Find the largest power of 2 that divides n - 1, call it 2^s.
	// Calculate d = (n - 1) / 2^s.
	one := big.NewInt(1)
	nm1 := new(big.Int).Sub(n, one) // n-1
	nm2 := new(big.Int).Sub(n, two) // n-2

	d := new(big.Int).Set(nm1) // n-1
	s := 0
	for d.Bit(0) == 0 {
		d.Rsh(d, 1) // d = (n - 1) / 2
		s++
	}

	for i := 0; i < k; i++ {
		a, err := rand.Int(rand.Reader, nm2) // Choose a random integer a in the range [2, n-2]
		if err != nil {
			panic(err)
		}

		x := new(big.Int).Exp(a, d, n) // Compute x = a^d mod n
		if x.Cmp(one) == 0 || x.Cmp(nm1) == 0 {
			continue // If x = 1 or x = n-1, (probably prime).
		}

		// Check for composite:
		for r := 1; r < s; r++ { // For i = 1 to s-1:
			x.Exp(x, two, n) // Compute x = x^2 mod n.
			if x.Cmp(nm1) == 0 {
				break // If x = n-1, (probably prime).
			}
		}

		if x.Cmp(nm1) != 0 {
			return false // If x != n-1, return False (composite).
		}
	}

	return true
}

func main() {
	num := new(big.Int).SetUint64(18446744073709551557)
	k := 10 // Number of iterations
	prime := isPrimeMillerRabin(num, k)

	if prime {
		fmt.Println(num, "is probably prime")
	} else {
		fmt.Println(num, "is not prime")
	}
}
