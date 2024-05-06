package formula_evaluator

import (
	"github.com/shopspring/decimal"
)

type (
	FormulaEvaluator interface {
		Evaluate(formula string, operandValues map[string]decimal.Decimal, operatorActions map[string]OperatorAction) (answer decimal.Decimal, err error)
	}

	OperatorAction func(a decimal.Decimal, b decimal.Decimal) decimal.Decimal
)
