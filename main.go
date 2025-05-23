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
	// valid := IsValidJson(string(data))
	// if valid == 1 {
	// 	check(errors.New("Invalid json file"))
	// }
	// fmt.Println(valid)
	// pairs := ExtractKeyValuePairs(string(data))
	// fmt.Println("Extracted key-value pairs:")
	// for k, v := range pairs {
	// 	fmt.Printf("  %q : %q\n", k, v)
	// }

	// parsed, validbool := ValidateAndExtractTyped(string(data))
	// if !validbool {
	// 	fmt.Println("❌ Invalid JSON")
	// 	return
	// }

	// fmt.Println("✅ Valid JSON with typed values:")
	// for k, v := range parsed {
	// 	fmt.Printf("  %q: %v (%T)\n", k, v, v)
	// }

	result, ok := ParseJSON(string(data))
	if !ok {
		fmt.Println("❌ Invalid JSON")
		os.Exit(1)
	}

	fmt.Println("✅ Valid JSON with typed values:")
	printValue(result, 0)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printValue(val interface{}, indent int) {
	prefix := strings.Repeat("  ", indent)
	switch v := val.(type) {
	case map[string]interface{}:
		fmt.Println(prefix + "{")
		for k, v2 := range v {
			fmt.Printf("%s  %q: ", prefix, k)
			printValue(v2, indent+1)
		}
		fmt.Println(prefix + "}")
	case []interface{}:
		fmt.Println(prefix + "[")
		for _, v2 := range v {
			printValue(v2, indent+1)
		}
		fmt.Println(prefix + "]")
	default:
		fmt.Printf("%v (%T)\n", v, v)
	}
}
