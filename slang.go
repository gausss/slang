package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: slang [file]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}

}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to slang")
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	run(text)
}

func runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	run(string(data))
}

func run(data string) {
	tokens := strings.Fields(data)

	for _, token := range tokens {
		fmt.Println(token)
	}
}
