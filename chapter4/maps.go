package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	characters := map[string]map[rune]int{
		"letter": make(map[rune]int),
		"number": make(map[rune]int),
		"symbol": make(map[rune]int),
		"space":  make(map[rune]int)}
	in := bufio.NewReader(os.Stdin)
	for {
		r, num, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Whoops!", err)
		}
		if r == unicode.ReplacementChar && num == 1 {
			continue
		}
		if unicode.IsLetter(r) {
			if characters["letter"] == nil {
				characters["letter"] = make(map[rune]int)
			}
			characters["letter"][r]++
		}
		if unicode.IsDigit(r) {
			if characters["digit"] == nil {
				characters["digit"] = make(map[rune]int)
			}
			characters["digit"][r]++
		}
		if unicode.IsPunct(r) {
			if characters["symbol"] == nil {
				characters["symbol"] = make(map[rune]int)
			}
			characters["symbol"][r]++
		}
		if unicode.IsSpace(r) {
			if characters["space"] == nil {
				characters["space"] = make(map[rune]int)
			}
			characters["space"][r]++
		}
	}
	for k, v := range characters {
		fmt.Println("Counting", k, ":")
		for a, b := range v {
			fmt.Println("\t", string(a), ":", b)
		}
	}
}
