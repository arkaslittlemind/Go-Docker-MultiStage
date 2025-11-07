package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

func calculator(op string, a, b float64) (float64, error) {
	switch op {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "division":
		if b == 0 {
			return 0, errors.New("cannot divide number by Zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation %q", op)
	}
}

func main() {
	op := flag.String("op", "add", "operation: add|subtraction|multiplication|division")
	a := flag.Float64("a", 0, "operand a")
	b := flag.Float64("b", 0, "operand b")
	flag.Parse()

	res, err := calculator(*op, *a, *b)
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(res)
}
