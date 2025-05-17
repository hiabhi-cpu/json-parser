package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func IsValidJson(data []byte) int {
	// fmt.Println(data)
	stk := stack.New()
	// stk.Push(10)
	if len(data) == 0 || data[0] != 123 {
		fmt.Println("Not valid Json")
		return 1
	}
	for _, val := range data {
		if val == '{' {
			stk.Push(rune(val))
		}
		if val == '}' {
			pop := stk.Pop()
			if r, ok := pop.(rune); ok {
				if val == '}' && r != '{' {
					fmt.Println("Not valid Json")
					return 1
				}
			} else {
				fmt.Println("Type assertion failed")
				return 1
			}
		}
	}
	if stk.Len() != 0 {
		fmt.Println("Not valid Json")
		return 1
	}
	fmt.Println("Valid Json")
	return 0
}
