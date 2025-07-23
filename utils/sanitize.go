package utils

import (
	"errors"
	"strings"
	"unicode"
)

// splitWordByCase splits a word like "barBaz" into ["bar", "Baz"]
func splitWordByCase(word string) []string {
	if len(word) == 0 {
		return nil
	}

	var result []string
	start := 0
	prev := rune(word[0])
	for i, r := range word[1:] {
		if unicode.IsUpper(r) && unicode.IsLower(prev) {
			result = append(result, word[start:i+1])
			start = i + 1
		}
		prev = r
	}
	result = append(result, word[start:])
	return result
}

// Sanitize transforms the input into camelCase or PascalCase,
// depending on the case of the first valid letter in the input.
// It removes all non-alphanumeric characters.
// It errors if the input is empty or the first valid character
// is a digit.
func Sanitize(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("command cannot be empty")
	}

	// Step 1: split input into chunks by non-alphanumeric
	rawWords := []string{}
	current := strings.Builder{}

	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			current.WriteRune(r)
		} else if current.Len() > 0 {
			rawWords = append(rawWords, current.String())
			current.Reset()
		}
	}
	if current.Len() > 0 {
		rawWords = append(rawWords, current.String())
	}

	if len(rawWords) == 0 {
		return "", errors.New("no valid characters in command name")
	}

	// Flatten compound words by splitting by case transitions
	var words []string
	for _, raw := range rawWords {
		words = append(words, splitWordByCase(raw)...)
	}

	firstRune := []rune(words[0])[0]
	if unicode.IsDigit(firstRune) {
		return "", errors.New("command name cannot start with a digit")
	}
	pascalCase := unicode.IsUpper(firstRune)

	// Build result
	var result strings.Builder
	for index, word := range words {
		if len(word) == 0 {
			continue
		}
		runes := []rune(word)
		first := unicode.ToUpper(runes[0])
		rest := string(runes[1:])
		if index == 0 && !pascalCase {
			first = unicode.ToLower(runes[0])
			rest = strings.ToLower(rest)
		} else {
			rest = strings.ToLower(rest)
		}
		result.WriteRune(first)
		result.WriteString(rest)
	}

	return result.String(), nil
}

// SanitizeStrict transforms the input into PascalCase or camelCase,
// using only letters [A-Za-z]. All other characters (including digits) are removed.
// It determines casing based on the first valid letter's original case.
// Returns error if no valid letters are found.
func SanitizeStrict(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("command cannot be empty")
	}

	// Step 1: split into letter-only word chunks
	rawWords := []string{}
	current := strings.Builder{}

	for _, r := range input {
		if unicode.IsLetter(r) {
			current.WriteRune(r)
		} else if current.Len() > 0 {
			rawWords = append(rawWords, current.String())
			current.Reset()
		}
	}
	if current.Len() > 0 {
		rawWords = append(rawWords, current.String())
	}

	if len(rawWords) == 0 {
		return "", errors.New("no valid characters in command name")
	}

	// Step 2: split each word by casing transitions
	var words []string
	for _, raw := range rawWords {
		words = append(words, splitWordByCase(raw)...)
	}

	// Step 3: determine casing from first letter
	firstRune := []rune(words[0])[0]
	pascalCase := unicode.IsUpper(firstRune)

	// Step 4: build result
	var result strings.Builder
	for index, word := range words {
		if len(word) == 0 {
			continue
		}
		runes := []rune(word)
		first := unicode.ToUpper(runes[0])
		rest := string(runes[1:])
		if index == 0 && !pascalCase {
			first = unicode.ToLower(runes[0])
			rest = strings.ToLower(rest)
		} else {
			rest = strings.ToLower(rest)
		}
		result.WriteRune(first)
		result.WriteString(rest)
	}

	return result.String(), nil
}
