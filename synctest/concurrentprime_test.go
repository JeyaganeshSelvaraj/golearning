package synctest

import (
	"fmt"
	"math"
	"testing"
	"time"
)

const LargestPrime = 100_000

func sieveOfEratosthenesSerial(n int) []int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}
	primeNumbers := []int{}
	for i := 2; i <= n; i++ {
		if primes[i] == true {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}
func generate(prime chan<- int) {
	fmt.Println("Prime", &prime)
	for i := 3; ; i += 2 {
		fmt.Println("Generating", &prime)
		prime <- i
	}
}
func filter(in <-chan int, out chan<- int, prime int) {
	fmt.Println("Prime1", &in)
	fmt.Println("Prime2", &out)
	fmt.Println("Filter for prime", prime)
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
func TestSerialPrime(t *testing.T) {
	start := time.Now()
	sieve := sieveOfEratosthenesSerial(LargestPrime)
	elapsed := time.Since(start)
	fmt.Println("\nComputation time: ", elapsed)
	fmt.Println("Largest prime: ", sieve[len(sieve)-1])
}

func TestConcurPrime(t *testing.T) {
	start := time.Now()
	prime1 := make(chan int)
	primes := []int{}
	go generate(prime1)
	for {
		prime := <-prime1
		if prime > LargestPrime {
			break
		}
		primes = append(primes, prime)
		prime2 := make(chan int)
		go filter(prime1, prime2, prime)
		prime1 = prime2
	}
	elapsed := time.Since(start)
	fmt.Println("Computation time: ", elapsed)
	fmt.Println("Number of primes = ", len(primes))
}

func primesBetween(prime []int, low, high int) []int {
	// Computes the prime numbers between low and high
	// given the initial set of primes from the SieveOfEratosthenes
	limit := high - low
	var result []int
	segment := make([]bool, limit+1)
	for i := 0; i < len(segment); i++ {
		segment[i] = true
	}
	// Find the primes in the current segment based on initial primes
	for i := 0; i < len(prime); i++ {
		lowlimit := int(math.Floor(float64(low)/float64(prime[i])) *
			float64(prime[i]))
		if lowlimit < low {
			lowlimit += prime[i]
		}
		for j := lowlimit; j < high; j += prime[i] {
			segment[j-low] = false
		}
	}
	for i := low; i < high; i++ {
		if segment[i-low] == true {
			result = append(result, i)
		}
	}
	return result
}
func SegmentedSieve(n int) []int {
	// Each segment is of size square root of n
	// Finds all primes up to n
	var primeNumbers []int
	limit := (int)(math.Floor(math.Sqrt(float64(n))))
	prime := sieveOfEratosthenesSerial(limit)
	for i := 0; i < len(prime); i++ {
		primeNumbers = append(primeNumbers, prime[i])
	}
	low := limit
	high := 2 * limit
	if high >= n {
		high = n
	}
	for {
		if low < n {
			next := primesBetween(prime, low, high)
			// fmt.Printf("\nprimesBetween(%d, %d) = %v", low, high, next)
			for i := 0; i < len(next); i++ {
				primeNumbers = append(primeNumbers, next[i])
			}
			low = low + limit
			high = high + limit
			if high >= n {
				high = n
			}
		} else {
			break
		}
	}
	return primeNumbers
}
func TestSegmentedSieve(t *testing.T) {
	start := time.Now()
	primeNumbers := SegmentedSieve(LargestPrime)
	elapsed := time.Since(start)
	fmt.Println("\nComputation time: ", elapsed)
	fmt.Println("Number of primes = ", len(primeNumbers))
}
