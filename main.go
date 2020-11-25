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

	if action == "uwu" {
		answer = uwuize(input)
	} else if action == "boomerspeak" {
		answer = boomerSpeak(input)
	} else {
		answer = "invalid cancer"
	}

	fmt.Println(answer)
}

func uwuize(input string) string {
	var output string
	strings.ReplaceAll(input, "r", "w")
	strings.ReplaceAll(input, "R", "W")
	strings.ReplaceAll(input, "l", "w")
	strings.ReplaceAll(input, "L", "W")

	for _, rune := range input {
		char := string(rune)
		outputLast := output[len(output)-1:]

		if (char == "o" || char == "O") && (strings.ContainsAny("MmNn", outputLast)) {
			output += "y"
		}
		output += string(rune)
	}

	return output
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
