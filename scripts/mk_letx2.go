//usr/local/go/bin/go run $0 $@ $(dirname `realpath $0`); exit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func loadWordsFromFile(filePath string) []string {
	fileDict, err := os.Open(filePath)
	assert(err)
	defer fileDict.Close()

	var words []string
	scanner := bufio.NewScanner(fileDict)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " ")

		if line != "" {
			words = append(words, line)
		}
	}

	return words
}

func main() {
	w1 := loadWordsFromFile("dict/letter2.txt")
	//fmt.Printf("Letter 2 count:%d\n", len(w1))

	w2 := loadWordsFromFile("dict/letter3.txt")
	//fmt.Printf("Letter 3 count:%d\n", len(w2))

	for i := 0; i < len(w1); i++ {
		for j := 0; j < len(w2); j++ {
			fmt.Printf("%s%s\n", w1[i], w2[j])
		}
	}
}
