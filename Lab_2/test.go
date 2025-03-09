package main

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConvertPrefixToLisp(t *testing.T) {
    cases := []struct {
        input    string
        expected string
        hasError bool
    }{
        {"+ 5 - 4 2 ^ 3 2", "(+ 5 (* (- 4 2) (pow 3 2)))", false},
        {"- 10 2", "(- 10 2)", false},
        {"* + 1 2 3", "(* (+ 1 2) 3)", false},
        {"", "", true},
        {"invalid input", "", true},
    }

    for _, tc := range cases {
        result, err := ConvertPrefixToLisp(tc.input)
        if tc.hasError {
            assert.Error(t, err)
        } else {
            assert.NoError(t, err)
            assert.Equal(t, tc.expected, result)
        }
    }
}

func ExampleConvertPrefixToLisp() {
    result, err := ConvertPrefixToLisp("+ 5 * - 4 2 ^ 3 2")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
    // Output: (+ 5 (* (- 4 2) (pow 3 2)))
}
