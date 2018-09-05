//usr/local/go/bin/go run $0 $@ $(dirname `realpath $0`); exit
package main

import "fmt"

func main() {
	for c := 'a'; c <= 'z'; c++ {
		for cc := 'a'; cc <= 'z'; cc++ {
			fmt.Printf("%c%ccn\n", c, cc)
		}
	}
}
