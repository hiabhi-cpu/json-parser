package main

import (
	"regexp"
	"strings"
)

func IsValidJson(json string) int {
	json = strings.TrimSpace(json)

	// Empty object is valid
	if json == "{}" {
		return 0
	}

	// Regex: {"key": "value", "key2": "value2", ...}
	pattern := `^\{\s*("(\\.|[^"\\])*"\s*:\s*"(\\.|[^"\\])*"\s*)(\s*,\s*"(\\.|[^"\\])*"\s*:\s*"(\\.|[^"\\])*"\s*)*\s*\}$`
	re := regexp.MustCompile(pattern)
	if re.MatchString(json) {
		return 0
	}
	return 1
}
