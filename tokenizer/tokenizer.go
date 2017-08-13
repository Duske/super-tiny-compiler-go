package tokenizer

import "regexp"

type Token struct {
	TypeName string
	Value string
}

func Build (input string) []Token {
	// track position in input string
	current := 0
	// length of the input
	inputLength := len(input)

	// we will collect all tokens in this slice
	tokens := []Token{}

	// iterate over each character
	for current < inputLength {
		// get the current character at current position
		char := input[current]

		// for opening parentheses add token
		if char == '(' {
			newToken := Token{TypeName: "paren", Value: "("}
			tokens = append(tokens, newToken)
			current++
			continue
		}

		// for closing parentheses add token
		if char == ')' {
			newToken := Token{TypeName: "paren", Value: ")"}
			tokens = append(tokens, newToken)
			current++
			continue
		}

		// if character is whitespace skip this position
		var reWhitespace = regexp.MustCompile(`\s`)
		if reWhitespace.MatchString(string(char)) {
			current++
			continue
		}

		// if character is a digit, concat all digits which directly follow this digit as a number
		var reDigit = regexp.MustCompile(`\d`)
		if reDigit.MatchString(string(char)) {
			value := ""
			for reDigit.MatchString(string(char)) {
				value += string(char)
				current += 1
				char = input[current]
			}
			newToken := Token{TypeName: "number", Value: value}
			tokens = append(tokens, newToken)
			continue
		}

		// if character is a string beginning with ",concat all following characters as a string
		if char == '"' {
			value := ""

			// skip " token
			current += 1
			char = input[current]

			for char != '"' {
				value += string(char)
				current += 1
				char = input[current]
			}

			// skip closing " token
			current += 1
			char = input[current]

			newToken := Token{TypeName: "string", Value: value}
			tokens = append(tokens, newToken)
			continue
		}

		// if character is a alphabetic char, add all following chars as a name
		var reAlphabet = regexp.MustCompile(`[a-z]`)
		if reAlphabet.MatchString(string(char)) {
			value := ""
			for reAlphabet.MatchString(string(char)) {
				value += string(char)
				current += 1
				char = input[current]
			}
			newToken := Token{TypeName: "name", Value: value}
			tokens = append(tokens, newToken)
			continue
		}
	}
	return tokens
}