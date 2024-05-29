package service

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"

	"github.com/Vojan-Najov/daec/pkg/rpn"
)

const (
	StatusError     = "Error"
	StatusDone      = "Done"
	StatusInProcess = "In process"
)

const (
	TokenTypeNumber = iota
	TokenTypeOperation
	TokenTypeTask
)

type Token interface {
	Type() int
}

type NumToken struct {
	Value float64
}

func (num NumToken) Type() int {
	return TokenTypeNumber
}

type OpToken struct {
	Value string
}

func (num OpToken) Type() int {
	return TokenTypeOperation
}

type TaskToken struct {
	ID int64
}

func (num TaskToken) Type() int {
	return TokenTypeTask
}

type Expression struct {
	*list.List
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

// Структура для ответа по запросу на endpoint expressions/{id}
type ExpressionUnit struct {
	Expr Expression `json:"expression`
}

// Структура для ответа по запросу на endpoint expressions
type ExpressionList struct {
	Exprs []Expression `json:"expressions"`
}

func NewExpression(id, expr string) (*Expression, error) {
	rpn, err := rpn.NewRPN(expr)
	fmt.Println(rpn)
	if err != nil {
		return nil, err
	}
	expression := Expression{
		List:   list.New(),
		ID:     id,
		Status: StatusInProcess,
		Result: "",
	}
	for _, val := range rpn {
		if strings.Contains("-+*/", val) {
			expression.PushBack(OpToken{val})
		} else {
			num, err := strconv.ParseFloat(val, 10)
			if err != nil {
				return nil, err
			}
			expression.PushBack(NumToken{num})
		}
	}
	fmt.Println(expression)
	return &expression, nil
}