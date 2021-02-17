package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	f("Direct")

	var tempSec = time.Now().Second()
	groupCount := 0

	for i := 0; i < 10000; i++ {
		func(msg string) {
			if time.Now().Second() != tempSec {
				fmt.Println("Total dispatched ", groupCount)
				time.Sleep(time.Second)
				groupCount = 0
				tempSec = time.Now().Second()
			} else {
				groupCount++
			}
		}("going")
	}

	time.Sleep(time.Second)
	fmt.Println("done")
}
