package main

import "fmt"

var hadError bool = false

func Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s \n", line, where, message)
	hadError = true
}
