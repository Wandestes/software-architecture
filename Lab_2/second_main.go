package main

import (
    "errors"
    "strings"
)

// ConvertPrefixToLisp converts a prefix expression to a Lisp-like expression
func ConvertPrefixToLisp(expression string) (string, error) {
    if expression == "" {
        return "", errors.New("empty expression")
    }
    tokens := strings.Fields(expression)
    index := 0
    result, err := parsePrefix(tokens, &index)
    if err != nil {
        return "", err
    }
    return result, nil
}

// parsePrefix recursively parses a prefix expression
func parsePrefix(tokens []string, index *int) (string, error) {
    if *index >= len(tokens) {
        return "", errors.New("invalid expression")
    }
    token := tokens[*index]
    *index++

    if isOperator(token) {
        left, err := parsePrefix(tokens, index)
        if err != nil {
            return "", err
        }
        right, err := parsePrefix(tokens, index)
        if err != nil {
            return "", err
        }
        if token == "^" {
            token = "pow"
        }
        return "(" + token + " " + left + " " + right + ")", nil
    }

    return token, nil
}

// isOperator checks if a token is a valid operator
func isOperator(token string) bool {
    switch token {
    case "+", "-", "*", "/", "^":
        return true
    }
    return false
}
