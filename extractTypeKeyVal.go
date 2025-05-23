package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ValidateAndExtractTyped(json string) (map[string]interface{}, bool) {
	json = strings.TrimSpace(json)
	if json == "{}" {
		return map[string]interface{}{}, true
	}

	key := `"((\\.|[^"\\])*)"`
	stringVal := `"((\\.|[^"\\])*)"`
	numberVal := `-?\d+(\.\d+)?([eE][+-]?\d+)?`
	boolVal := `true|false`
	nullVal := `null`
	value := fmt.Sprintf(`(%s|%s|%s|%s)`, stringVal, numberVal, boolVal, nullVal)

	pairRegex := regexp.MustCompile(fmt.Sprintf(`%s\s*:\s*%s`, key, value))
	matches := pairRegex.FindAllStringSubmatch(json, -1)
	if matches == nil {
		return nil, false
	}

	result := make(map[string]interface{})

	for _, match := range matches {
		key := unescape(match[1])
		rawVal := match[3] // this group gets the actual value part (string, number, etc)

		if strings.HasPrefix(rawVal, `"`) && strings.HasSuffix(rawVal, `"`) {
			// It's a string
			unquoted := rawVal[1 : len(rawVal)-1]
			result[key] = unescape(unquoted)
		} else if rawVal == "true" {
			result[key] = true
		} else if rawVal == "false" {
			result[key] = false
		} else if rawVal == "null" {
			result[key] = nil
		} else {
			// Try parse as float
			num, err := strconv.ParseFloat(rawVal, 64)
			if err != nil {
				return nil, false
			}
			result[key] = num
		}
	}

	// Basic check for structure integrity
	content := strings.Trim(json, "{} \n\t")
	if len(content) > 0 && strings.Count(content, ",")+1 != len(matches) {
		return nil, false
	}

	return result, true
}

func unescape(s string) string {
	return strings.NewReplacer(`\"`, `"`, `\\`, `\`).Replace(s)
}
