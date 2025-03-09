package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)


type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}


func (h *ComputeHandler) Compute() error {
	inputData, err := io.ReadAll(h.Input)
	if err != nil {
		return err
	}

	expression := strings.TrimSpace(string(inputData))
	if expression == "" {
		return fmt.Errorf("empty expression")
	}


	result, err := ConvertPrefixToLisp(expression)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(h.Output, result)
	return err
}

func main() {
	exprFlag := flag.String("e", "", "Expression in prefix notation")
	fileFlag := flag.String("f", "", "File containing expression")
	outFileFlag := flag.String("o", "", "Output file")

	flag.Parse()


	if *exprFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: specify either -e or -f, not both")
		os.Exit(1)
	}


	var input io.Reader
	if *exprFlag != "" {
		input = strings.NewReader(*exprFlag)
	} else if *fileFlag != "" {
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Error: no input provided")
		os.Exit(1)
	}


	var output io.Writer = os.Stdout
	if *outFileFlag != "" {
		file, err := os.Create(*outFileFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	handler := ComputeHandler{Input: input, Output: output}
	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error processing expression:", err)
		os.Exit(1)
	}
}

