package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler_Success(t *testing.T) {
	input := strings.NewReader("+ 5 * - 4 2 ^ 3 2")
	var output strings.Builder

	handler := ComputeHandler{Input: input, Output: &output}
	err := handler.Compute()

	assert.NoError(t, err)
	assert.Equal(t, "(+ 5 (* (- 4 2) (pow 3 2)))\n", output.String())
}

func TestComputeHandler_EmptyInput(t *testing.T) {
	input := strings.NewReader("")
	var output strings.Builder

	handler := ComputeHandler{Input: input, Output: &output}
	err := handler.Compute()

	assert.Error(t, err)
	assert.Empty(t, output.String())
}

func TestComputeHandler_InvalidExpression(t *testing.T) {
	input := strings.NewReader("invalid input")
	var output strings.Builder

	handler := ComputeHandler{Input: input, Output: &output}
	err := handler.Compute()

	assert.Error(t, err)
	assert.Empty(t, output.String())
}
