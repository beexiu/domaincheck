package main

import "fmt"

func main() {
	for c := 'a'; c <= 'z'; c++ {
		for cc := 'a'; cc <= 'z'; cc++ {
			fmt.Printf("at%c%c\n", c, cc)
		}
	}
}
