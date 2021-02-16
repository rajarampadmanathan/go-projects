package main

import "fmt"

func oddeven() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
