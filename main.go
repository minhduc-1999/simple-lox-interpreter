package main

import (
	"bufio"
	"fmt"
	"lox/lexer"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: lox [script]")
		os.Exit(1)
	}
	if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		data := scanner.Text()
		if data == ".exit" {
			os.Exit(0)
		}
		run(data)
		fmt.Print("> ")
	}
}

func runFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to read file %v", path))
		os.Exit(1)
	}
	run(string(data))
}

func run(s string) {
	lexer := lexer.NewLexer(s)
	tokens, errors := lexer.ScanToken()
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		for _, token := range tokens {
			fmt.Println(fmt.Sprintf("%#v", token))
		}
	}
}
