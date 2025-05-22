package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello")
	if len(os.Args) == 1 {
		check(errors.New("Give a json file name usage:\n json-parser <FILE_NAME>"))
	}
	if !strings.Contains(os.Args[1], ".json") {
		check(errors.New("Give correct json file"))
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		check(err)
	}
	valid := IsValidJson(string(data))
	if valid == 1 {
		check(errors.New("Invalid json file"))
	}
	fmt.Println(valid)
	pairs := ExtractKeyValuePairs(string(data))
	fmt.Println("Extracted key-value pairs:")
	for k, v := range pairs {
		fmt.Printf("  %q : %q\n", k, v)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
