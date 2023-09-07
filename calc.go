package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculator App")
	fmt.Println("Enter a calculation (Example: 1 + 2 or 2 * 5), or type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		fmt.Printf("Result: %d\n", result)
	}
}

func calculate(input string) (int, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, fmt.Errorf("Invalid input")
	}

	left, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	right, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	operator := parts[1]
	switch operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("Division by zero is not allowed")
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("Invalid operator: %s", operator)
	}
}
