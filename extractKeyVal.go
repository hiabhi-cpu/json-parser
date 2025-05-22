package main

import (
	"regexp"
	"strings"
)

func ExtractKeyValuePairs(json string) map[string]string {
	pairs := make(map[string]string)

	// Regex to extract key-value pairs
	pairRegex := regexp.MustCompile(`"((?:\\.|[^"\\])*)"\s*:\s*"((?:\\.|[^"\\])*)"`)

	matches := pairRegex.FindAllStringSubmatch(json, -1)
	for _, match := range matches {
		key := unescapeString(match[1])
		value := unescapeString(match[2])
		pairs[key] = value
	}
	return pairs
}

func unescapeString(s string) string {
	replacer := strings.NewReplacer(
		`\"`, `"`,
		`\\`, `\`,
	)
	return replacer.Replace(s)
}
