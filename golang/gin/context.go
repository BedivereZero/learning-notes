package main

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

type Operator struct {
	Name string
}

type contextKey struct{}

func NewContextWithOperator(parent context.Context, o *Operator) context.Context {
	return context.WithValue(parent, contextKey{}, o)
}

func OperatorFromContext(ctx context.Context) (*Operator, error) {
	v := ctx.Value(contextKey{})
	if v == nil {
		return nil, errors.New("operator is not found")
	}

	o, ok := v.(*Operator)
	if !ok {
		return nil, fmt.Errorf("unexpected data type for context key: %T", v)
	}

	return o, nil
}

func WithOperator(operators ...Operator) gin.HandlerFunc {
	if operators == nil {
		operators = append(operators, Operator{Name: "Example"})
	}

	var index atomic.Int64

	return func(c *gin.Context) {
		i := int(index.Load()) % len(operators)
		c.Request = c.Request.WithContext(NewContextWithOperator(c.Request.Context(), &operators[i]))
	}
}
