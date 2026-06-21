package main

import "fmt"

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("non dividable")
	} else {
		return a / b, nil
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for tmp := 2; tmp < n; tmp++ {
		if n % tmp == 0 {
			return false
		}
	}
	return true
}

func main() {
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	for tmp := 0; tmp <= 50; tmp++ {
		check := isPrime(tmp)
		if check {
			fmt.Printf("%d is prime\n", tmp)
		} else {
			fmt.Printf("%d isn't prime\n", tmp)
		}
	}
}
