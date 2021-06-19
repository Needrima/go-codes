//Sieve of erastothenes
//Generating prime numbers
package main

import (
	"fmt"
)

func Generate(in chan uint64) { //Generate integers and send to in
	for i := uint64(2); ; i++ {
		in <- i
	}
}

/*Receive values from in and send those not divisible by
prime to out
*/
func Filter(in, out chan uint64, prime uint64) {
	for {
		v := <-in
		if v%prime != 0 {
			out <- v
		}
	}
}

func main() {
	in := make(chan uint64) //declare channel in
	go Generate(in)         //launch number generators

	for i := 0; i < 1000000; i++ {
		prime := <-in
		fmt.Println(prime)
		out := make(chan uint64)
		go Filter(in, out, prime)
		in = out
	}
}
