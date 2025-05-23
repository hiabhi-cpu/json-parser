package main

import (
	"strconv"
	"strings"
)

func ParseJSON(input string) (interface{}, bool) {
	input = strings.TrimSpace(input)

	switch {
	case strings.HasPrefix(input, "{") && strings.HasSuffix(input, "}"):
		return parseObject(input)
	case strings.HasPrefix(input, "[") && strings.HasSuffix(input, "]"):
		return parseArray(input)
	case input == "true":
		return true, true
	case input == "false":
		return false, true
	case input == "null":
		return nil, true
	case strings.HasPrefix(input, `"`) && strings.HasSuffix(input, `"`):
		return unescape(input[1 : len(input)-1]), true
	default:
		num, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return num, true
		}
	}

	return nil, false
}

func parseObject(input string) (map[string]interface{}, bool) {
	result := make(map[string]interface{})
	content := strings.Trim(input, "{} \n\t")

	if content == "" {
		return result, true
	}

	pairs, ok := splitTopLevel(content, ',')
	if !ok {
		return nil, false
	}

	for _, pair := range pairs {
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) != 2 {
			return nil, false
		}
		key := strings.TrimSpace(kv[0])
		val := strings.TrimSpace(kv[1])

		if !strings.HasPrefix(key, `"`) || !strings.HasSuffix(key, `"`) {
			return nil, false
		}
		key = unescape(key[1 : len(key)-1])

		parsedVal, ok := ParseJSON(val)
		if !ok {
			return nil, false
		}

		result[key] = parsedVal
	}

	return result, true
}

func parseArray(input string) ([]interface{}, bool) {
	content := strings.Trim(input, "[] \n\t")
	if content == "" {
		return []interface{}{}, true
	}

	elems, ok := splitTopLevel(content, ',')
	if !ok {
		return nil, false
	}

	var result []interface{}
	for _, elem := range elems {
		elem = strings.TrimSpace(elem)
		val, ok := ParseJSON(elem)
		if !ok {
			return nil, false
		}
		result = append(result, val)
	}
	return result, true
}

// splitTopLevel splits a string by a delimiter at top level, ignoring nested brackets
func splitTopLevel(s string, delim rune) ([]string, bool) {
	var parts []string
	var buf strings.Builder
	depth := 0
	inString := false

	for i, r := range s {
		switch r {
		case '"':
			if i == 0 || s[i-1] != '\\' {
				inString = !inString
			}
		case '{', '[':
			if !inString {
				depth++
			}
		case '}', ']':
			if !inString {
				depth--
			}
		case delim:
			if depth == 0 && !inString {
				parts = append(parts, strings.TrimSpace(buf.String()))
				buf.Reset()
				continue
			}
		}
		buf.WriteRune(r)
	}

	if depth != 0 || inString {
		return nil, false
	}

	parts = append(parts, strings.TrimSpace(buf.String()))
	return parts, true
}
