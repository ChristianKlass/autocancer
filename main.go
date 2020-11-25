package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("It's -> autocancer.go <cancertype> <input>")
		return
	}

	var answer string
	action := os.Args[1]
	input := os.Args[2]

	if action == "cancerize" {
		answer = cancerize(input)
	} else if action == "boomerspeak" {
		answer = boomerSpeak(input)
	} else {
		answer = "invalid cancer"
	}

	fmt.Println(answer)
}

func cancerize(input string) string {
	input += " cancered"
	return input
}

func boomerSpeak(input string) string {
	var output string

	input = strings.ToLower(input)
	for index, rune := range input {
		if index%2 == 1 {
			output += strings.ToUpper(string(rune))
		} else {
			output += string(rune)
		}
	}

	return output
}
