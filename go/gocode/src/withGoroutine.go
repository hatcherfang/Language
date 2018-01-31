// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	go f("goroutine")
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.

	// You can also start a goroutine for an anonymous
	// function call.
	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.
	// time.Sleep(time.Second)
	// var input string
	// fmt.Scanln(&input)
	// fmt.Println("done")
}
