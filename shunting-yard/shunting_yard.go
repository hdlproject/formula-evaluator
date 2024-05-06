package shunting_yard

import (
	"fmt"
)

var operatorPrecedenceMap = map[string]int{
	"+": 2,
	"-": 2,
	"x": 3,
	"/": 3,
	"^": 4,
}

type operatorAssociativity int

const (
	operatorAssociativityLeft = iota
	operatorAssociativityRight
)

var operatorAssociativityMap = map[string]operatorAssociativity{
	"+": operatorAssociativityLeft,
	"-": operatorAssociativityLeft,
	"x": operatorAssociativityLeft,
	"/": operatorAssociativityLeft,
	"^": operatorAssociativityRight,
}

type shuntingYard struct {
	outputQueue []string
}

//func (s *shuntingYard) Evaluate(formula string,
//	operandValues map[string]decimal.Decimal,
//	operatorActions map[string]formula_evaluator.OperatorAction,
//) (answer decimal.Decimal, err error) {
//	answer = decimal.Zero
//
//	postfixFormula := getPostfixFormula(formula)
//
//	return answer, nil
//}

func isNumber(token rune) bool {
	return token >= 48 && token <= 57
}

func isLeftParentheses(token string) bool {
	return token == "("
}

func isRightParentheses(token string) bool {
	return token == ")"
}

func getPostfixFormula(formula string) (postfixFormula []string) {
	var operatorStack stringStack

	for _, token := range formula {
		tokenStr := fmt.Sprintf("%s", string(token))

		if isNumber(token) {
			postfixFormula = append(postfixFormula, tokenStr)
		} else {
			topStack, isFound := operatorStack.Pop()
			if !isFound {
				operatorStack.Push(tokenStr)
				continue
			}

			if !isRightParentheses(tokenStr) {
				for isFound {
					if !isLeftParentheses(tokenStr) &&
						operatorPrecedenceMap[topStack] >= operatorPrecedenceMap[tokenStr] &&
						operatorAssociativityMap[tokenStr] == operatorAssociativityLeft {

						postfixFormula = append(postfixFormula, topStack)
					} else {
						operatorStack.Push(topStack)
						break
					}

					topStack, isFound = operatorStack.Pop()
				}
				operatorStack.Push(tokenStr)
			} else {
				for isFound {
					if isLeftParentheses(topStack) {
						break
					}

					postfixFormula = append(postfixFormula, topStack)

					topStack, isFound = operatorStack.Pop()
				}
			}
		}
	}

	topStack, isFound := operatorStack.Pop()
	for isFound {
		postfixFormula = append(postfixFormula, topStack)

		topStack, isFound = operatorStack.Pop()
	}

	return postfixFormula
}
