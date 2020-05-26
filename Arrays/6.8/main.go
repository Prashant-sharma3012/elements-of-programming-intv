package main

import "fmt"

// list all prime numbers between 1 to input
func main() {
	// input is n a number
	n := 43
	var inputArr []int
	var isPrime []bool
	var primes []int

	for i := 0; i <= 43; i++ {
		inputArr = append(inputArr, i)

		if i == 0 || i == 1 {
			isPrime = append(isPrime, false)
		} else {
			isPrime = append(isPrime, true)
		}
	}

	for indx, val := range inputArr {
		if isPrime[indx] {
			primes = append(primes, val)

			for j := indx + 1; j <= n; j++ {
				if j%val == 0 && j < len(isPrime) {
					isPrime[j] = false
				}
			}
		}
	}

	fmt.Println(primes)
}
