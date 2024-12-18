package calculation

import (
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if !isValid(expression) {
		return 0, fmt.Errorf("invalid expression")
	}
	postfix := infixToPostfix(expression)

	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func isValid(e string) (bool) {
	res := []rune(e)
	if len(res) == 0 {
        return false
    }
	if res[0] == '+' || res[0] == '-' || res[0] == '*' || res[0] == '/' {
		return false
	}
	var stack []rune
	for i, sym := range res {
		if  sym == '+' || sym == '-' || sym == '*' || sym == '/' {
			continue
		} else if sym == '1' || sym == '2' || sym == '3' || sym == '4' || sym == '5' || sym == '6' || sym == '7' || sym == '8' || sym == '9' || sym == '0' {
			if i + 1 != len(res) {
				if !(res[i + 1] == '1' || res[i + 1] == '2' || res[i + 1] == '3' || res[i + 1] == '4' || res[i + 1] == '5' || res[i + 1] == '6' || res[i + 1] == '7' || res[i + 1] == '8' || res[i + 1] == '9' || res[i + 1] == '0' || res[i + 1] == '+' || res[i + 1] == '-' || res[i + 1] == '*' || res[i + 1] == '/' || res[i + 1] == ')') {
					return false
				}
			}
			if i - 1 != -1 {
				if !(res[i - 1] == '1' || res[i - 1] == '2' || res[i - 1] == '3' || res[i - 1] == '4' || res[i - 1] == '5' || res[i - 1] == '6' || res[i - 1] == '7' || res[i - 1] == '8' || res[i - 1] == '9' || res[i - 1] == '0' || res[i - 1] == '+' || res[i - 1] == '-' || res[i - 1] == '*' || res[i - 1] == '/' || res[i - 1] == '(') {
					return false
				} 
			} 
		} else if sym == '(' {
			stack = append(stack, sym)
		} else if sym == ')' {
			if stack == nil {
				stack = append(stack, sym)
			} else if stack[len(stack) - 1] == '(' {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if stack == nil {
		return true
	} else if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func infixToPostfix(expression string) []string {
	var postfix []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	tokens := strings.Split(expression, "")

	for _, token := range tokens {
		if token == "(" {
			stack = append(stack, token)
		} else if token == ")" {
			for stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else if _, isOperator := precedence[token]; isOperator {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else {
			postfix = append(postfix, token)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func evaluatePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}

			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}