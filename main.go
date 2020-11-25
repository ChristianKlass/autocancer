package main

import (
	"fmt"
	"os"
)

func main() {
	action := os.Args[1]
	input := os.Args[2]
	var answer string

	if input != "" {
		if action == "cancerize" {
			answer = cancerize(input)
		} else if action == "boomerspeak" {
			answer = boomerSpeak(input)
		} else {
			answer = "invalid cancer"
		}
	} else {
		fmt.Println("There's no input to cancerize!")
	}

	fmt.Println(answer)
}

func cancerize(input string) string {
	input += " cancered"
	return input
}

func boomerSpeak(input string) string {
	input += " boomered"
	return input
}
