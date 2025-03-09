package main

import (
	"errors"
	"strings"
)

// ConvertPrefixToLisp converts a prefix expression to Lisp-like notation.
func ConvertPrefixToLisp(expression string) (string, error) {
	if expression == "" {
		return "", errors.New("empty expression")
	}

	// Split the expression into tokens
	tokens := strings.Fields(expression)
	index := 0

	// Parse the tokens recursively
	result, err := parsePrefix(tokens, &index)
	if err != nil {
		return "", err
	}

	// Ensure the parsed result does not have extra prefixes or issues.
	return result, nil
}

// parsePrefix recursively parses a prefix expression.
func parsePrefix(tokens []string, index *int) (string, error) {
	// If we've reached the end of the token list, return an error
	if *index >= len(tokens) {
		return "", errors.New("invalid expression")
	}

	// Get the current token and increment the index
	token := tokens[*index]
	*index++

	// If the token is an operator, parse the left and right operands recursively
	if isOperator(token) {
		left, err := parsePrefix(tokens, index)
		if err != nil {
			return "", err
		}
		right, err := parsePrefix(tokens, index)
		if err != nil {
			return "", err
		}

		// Handle the exponentiation operator (^), convert it to "pow"
		if token == "^" {
			token = "pow"
		}

		// Format the result as a Lisp expression
		return "(" + token + " " + left + " " + right + ")", nil
	}

	// If the token is not an operator, return it as a string (operand)
	return token, nil
}

// isOperator checks if the token is a valid operator.
func isOperator(token string) bool {
	switch token {
	case "+", "-", "*", "/", "^":
		return true
	}
	return false
}
