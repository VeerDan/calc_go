package calculation

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	err := isValid(expression)
	if err != nil {
		return 0, err
	}
	postfix := infixToPostfix(expression)
	result, err := evaluatePostfix(postfix)
	if err != nil {
		return 0, err
	}
	if math.IsInf(result, 0) {
		return 0, ErrDivisionByZero
	}
	return result, nil
}

func isValid(e string) error {
	res := []rune(strings.ReplaceAll(e, " ", ""))
	//проверка на пустую строку
	if len(res) == 0 {
		return ErrEmptyExpression
	}
	//Проверка первого и последнего символа на корректность
	if res[0] == ')' || res[0] == '+' || res[0] == '-' || res[0] == '*' || res[0] == '/' || res[len(res)-1] == '+' || res[len(res)-1] == '-' || res[len(res)-1] == '*' || res[len(res)-1] == '/' || res[len(res) - 1] == '(' {
		return ErrInvalidExpression
	}
	//Проверка, что первый символ либо цифра, либо `(`
	if !(res[0] == '1' || res[0] == '2' || res[0] == '3' || res[0] == '4' || res[0] == '5' || res[0] == '6' || res[0] == '7' || res[0] == '8' || res[0] == '9' || res[0] == '0' || res[0] == '(') {
		return ErrInvalidSymbol
	}
	//стэк для `(` и `)`
	var stack []rune
	//проверим правильность выражения посимвольно
	for i, sym := range res {
		//если знак операция - проверяем, что вокруг цифры
		if sym == '+' || sym == '-' || sym == '*' || sym == '/' {
			if !(res[i-1] == '1' || res[i-1] == '2' || res[i-1] == '3' || res[i-1] == '4' || res[i-1] == '5' || res[i-1] == '6' || res[i-1] == '7' || res[i-1] == '8' || res[i-1] == '9' || res[i-1] == '0' || res[i-1] == ')' || res[i+1] == '1' || res[i+1] == '2' || res[i+1] == '3' || res[i+1] == '4' || res[i+1] == '5' || res[i+1] == '6' || res[i+1] == '7' || res[i+1] == '8' || res[i+1] == '9' || res[i+1] == '0' || res[i+1] == '(') {
				return ErrInvalidExpression
			}
		} else if sym == '1' || sym == '2' || sym == '3' || sym == '4' || sym == '5' || sym == '6' || sym == '7' || sym == '8' || sym == '9' || sym == '0' {
			if i+1 != len(res) {
				if !(res[i+1] == '1' || res[i+1] == '2' || res[i+1] == '3' || res[i+1] == '4' || res[i+1] == '5' || res[i+1] == '6' || res[i+1] == '7' || res[i+1] == '8' || res[i+1] == '9' || res[i+1] == '0' || res[i+1] == '+' || res[i+1] == '-' || res[i+1] == '*' || res[i+1] == '/' || res[i+1] == ')') {
					//после цифры не может стоять `(`
					if res[i+1] == '(' {
						return ErrInvalidExpression
					}
					//в ином случае - посторонний символ
					return ErrInvalidSymbol
				}
			}
			if i-1 != -1 {
				if !(res[i-1] == '1' || res[i-1] == '2' || res[i-1] == '3' || res[i-1] == '4' || res[i-1] == '5' || res[i-1] == '6' || res[i-1] == '7' || res[i-1] == '8' || res[i-1] == '9' || res[i-1] == '0' || res[i-1] == '+' || res[i-1] == '-' || res[i-1] == '*' || res[i-1] == '/' || res[i-1] == '(') {
					if res[i-1] == ')' {
						return ErrInvalidExpression
					}
					return ErrInvalidSymbol
				}
			}
		} else if sym == '(' { //добавляем `(` в стэк проверки
			stack = append(stack, sym)
		} else if sym == ')' { //проверяем на наличие открывающих скобок
			if stack == nil {
				//нет открывающих скобок до закрывающих
				return ErrInvalidExpression
			} else if stack[len(stack)-1] == '(' {
				//если последний элемент открывающая скобка - убираем её из стэка
				stack = stack[:len(stack)-1]
			} else {
				return ErrInvalidExpression
			}
		} else {
			//посторонний символ
			return ErrInvalidSymbol
		}
	}

	if stack == nil {
		//скобок не было, ошибок нет
		return nil
	} else if len(stack) == 0 {
		return nil
	} else {
		return ErrInvalidExpression
	}
}

func infixToPostfix(expression string) []string {
	var postfix []string
	var stack []string
	var temp string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}

	tokens := strings.Split(expression, "")
	for i, token := range tokens {
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
			if (i + 1) < len(tokens) {
				if tokens[i + 1] == "1" || tokens[i + 1] == "2" || tokens[i + 1] == "3" || tokens[i + 1] == "4" || tokens[i + 1] == "5" || tokens[i + 1] == "6" || tokens[i + 1] == "7" || tokens[i + 1] == "8" || tokens[i + 1] == "9" || tokens[i + 1] == "0" {
					temp += token
				} else {
					temp += token
					postfix = append(postfix, temp)
					temp = ""
				}
			} else {
				temp += token
				postfix = append(postfix, temp)
				temp = ""
			}
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
				fmt.Println("1")
				return 0, ErrInvalidExpression
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
		return 0, ErrInvalidExpression
	}

	return stack[0], nil
}
