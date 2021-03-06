// In 6.6 we saw a recursive solution for calculating Fibonacci numbers. But
// they can also be calculated in an imperative way, using a simple array. Do 
// this for the first 50 Fibonacci numbers.
package main

import (
	"fmt"
	//"errors"
	//"os"
)

const (
	ARRSIZE int = 51
)

var fibs [ARRSIZE]uint64

/* memoization helps if fibo is called more than once, but no improvement the 
first time fibo(n) is computed */

func main() {
	for i := 0; i < 51; i++ {
		fibn, err := fibo(i)
		if err != nil {
			fmt.Printf("%v\n", err)
			break
		}
		fmt.Printf("Fib(%d): %d\n", i, fibn)
	}
}

func fibo(n int) (fibn uint64, err error) {
	var arr = [3]uint64{0, 1, 1}
	err = nil
	if n == 0 {
		return uint64(n), err
	}
	if n >= ARRSIZE {
		return 0, fmt.Errorf("Array overflow: n = %d Cannot be >= %v", n, ARRSIZE)
	}
	if fibs[n] != 0 {
		fibn = fibs[n]
		return fibn, err
	}
	for i := 2; i < n; i++ {
		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[0] + arr[1]
	}
	fibn = arr[2]
	fibs[n] = fibn
	return fibn, err
}
