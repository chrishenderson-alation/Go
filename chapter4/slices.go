package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func reverse(array *[5]int) *[5]int {
	i := 0
	j := len(*array) - 1
	for i < j {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
		i++
		j--
	}
	return array
}

func deduplicate(strings []string) []string {
	var index int
	var previous string
	for index < len(strings) {
		if strings[index] == previous {
			strings = append(strings[:index], strings[index+1:]...)
		} else {
			previous = strings[index]
			index++
		}
	}
	return strings
}

func squash(bytes []byte) []byte {
	var index int
	consuming := false
	for index < len(bytes) {
		if !consuming {
			if unicode.IsSpace(rune(bytes[index])) {
				consuming = true
			}
			index++
		} else if unicode.IsSpace(rune(bytes[index])) {
			bytes = append(bytes[:index], bytes[index+1:]...)
		} else {
			index++
			consuming = false
		}
	}
	return bytes
}

func squash2(things []byte) []byte {
	var index int
	prefix := []byte{' ', ' '}
	for index < len(things) {
		if bytes.HasPrefix(things[index:], prefix) {
			things = append(things[:index], things[index+1:]...)
		} else {
			index++
		}
	}
	return things
}

func main() {
	stuff := [...]int{1, 2, 3, 4, 5}
	reverse(&stuff)
	fmt.Println(stuff)
	things := []string{"a", "b", "b", "b", "c", "c", "a", "a"}
	things = deduplicate(things)
	fmt.Println(things)
	fmt.Println(squash([]byte{' ', ' ', 'a', ' ', ' ', ' '}))
	fmt.Println(squash2([]byte{' ', ' ', 'a', ' ', ' ', ' '}))
}
