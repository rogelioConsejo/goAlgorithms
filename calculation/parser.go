package calculation

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseAndCalculate parses a mathematical expression represented as a string and returns the computed integer result.
// The expression can contain parentheses, addition, subtraction, multiplication, and division.
// It uses recursive descent parsing to build an abstract syntax tree (AST) of the expression, and evaluates the AST.
// The expression can contain positive and negative integers.
func ParseAndCalculate(s string) int {
	calc := NewCalculation(s)
	return calc.Root().Value()
}

type Calculation interface {
	Root() Node
}

func NewCalculation(s string) Calculation {
	calc := parseCalculation(s)

	return calc
}

func parseCalculation(s string) calculation {
	if s == "" {
		return calculation{
			root: nilNode{},
		}
	}
	if parsedInt, err := strconv.ParseInt(s, 10, 64); err == nil {
		return calculation{
			root: numericValueNode{
				value: int(parsedInt),
			},
		}
	}

	if hasParentheses := strings.Contains(s, "("); hasParentheses {
		equivalentCalculation := replaceLastParenthesisWithValue(s)
		return parseCalculation(equivalentCalculation)
	}
	if hasPlus := strings.Contains(s, "+"); hasPlus {
		return parseSum(s)
	}
	if hasMinus := strings.Contains(s, "-"); hasMinus {
		return parseSubtraction(s)
	}
	if hasMultiplication := strings.Contains(s, "*"); hasMultiplication {
		return parseMultiplication(s)
	}
	if hasDivision := strings.Contains(s, "/"); hasDivision {
		return parseDivision(s)
	}
	return calculation{
		root: nilNode{},
	}
}

func parseDivision(s string) calculation {
	divTerms := strings.Split(s, "/")
	leftTerm := parseCalculation(divTerms[0])
	rightTerm := parseCalculation(strings.Join(divTerms[1:], "/"))
	calc := &calculation{}
	calc.root = &operatorNode{
		operation: func(a, b int) int {
			return a / b
		},
		left: &numericValueNode{
			value: leftTerm.root.Value(),
		},
		right: &numericValueNode{
			value: rightTerm.root.Value(),
		},
	}
	return *calc
}

func parseMultiplication(s string) calculation {
	multTerms := strings.Split(s, "*")
	leftTerm := parseCalculation(multTerms[0])
	rightTerm := parseCalculation(strings.Join(multTerms[1:], "*"))
	calc := &calculation{}
	calc.root = &operatorNode{
		operation: func(a, b int) int {
			return a * b
		},
		left: &numericValueNode{
			value: leftTerm.root.Value(),
		},
		right: &numericValueNode{
			value: rightTerm.root.Value(),
		},
	}
	return *calc
}

func parseSubtraction(s string) calculation {
	sumTerms := strings.Split(s, "-")
	leftTerm := parseCalculation(strings.Join(sumTerms[:len(sumTerms)-1], "-"))
	rightTerm := parseCalculation(sumTerms[len(sumTerms)-1])
	calc := &calculation{}
	calc.root = &operatorNode{
		operation: func(a, b int) int {
			return a - b
		},
		left: &numericValueNode{
			value: leftTerm.root.Value(),
		},
		right: &numericValueNode{
			value: rightTerm.root.Value(),
		},
	}
	return *calc
}

func parseSum(s string) calculation {
	sumTerms := strings.Split(s, "+")
	leftTerm := parseCalculation(sumTerms[0])

	rightTerm := parseCalculation(strings.Join(sumTerms[1:], "+"))

	calc := &calculation{}
	calc.root = &operatorNode{
		operation: func(a, b int) int {
			return a + b
		},
		left: &numericValueNode{
			value: leftTerm.root.Value(),
		},
		right: &numericValueNode{
			value: rightTerm.root.Value(),
		},
	}
	return *calc
}

func replaceLastParenthesisWithValue(s string) string {
	lastOpenParentheses := strings.LastIndex(s, "(")
	lastCloseParentheses := lastOpenParentheses + strings.Index(s[lastOpenParentheses:], ")")
	innerCalc := parseCalculation(s[lastOpenParentheses+1 : lastCloseParentheses])
	equivalentCalculation := s[:lastOpenParentheses] + fmt.Sprintf("%d", innerCalc.root.Value()) + s[lastCloseParentheses+1:]
	return equivalentCalculation
}

type calculation struct {
	root Node
}

func (c calculation) Root() Node {
	return c.root
}

type Node interface {
	Value() int
	Left() Node
	Right() Node
}

type numericValueNode struct {
	value int
}

func (n numericValueNode) Value() int {
	return n.value
}

func (n numericValueNode) Left() Node {
	return nilNode{}
}

func (n numericValueNode) Right() Node {
	return nilNode{}
}

type nilNode struct {
}

func (n nilNode) Value() int {
	return 0
}

func (n nilNode) Left() Node {
	return nilNode{}
}

func (n nilNode) Right() Node {
	return nilNode{}
}

type operatorNode struct {
	operation func(int, int) int
	left      Node
	right     Node
}

func (o operatorNode) Value() int {
	return o.operation(o.left.Value(), o.right.Value())
}

func (o operatorNode) Left() Node {
	return o.left
}

func (o operatorNode) Right() Node {
	return o.right
}
