package domain

import (
	"fmt"
)

type (
	AccountID struct {
		value string
	}
)

func NewAccountID(value string) AccountID {
	return AccountID{value: value}
}

func (a AccountID) AggregateIDTypeName() string {
	return "account"
}

func (a AccountID) Value() string {
	return a.value
}

func (a AccountID) AsString() string {
	return fmt.Sprintf("%s-%s", a.AggregateIDTypeName(), a.Value())
}

func (a AccountID) Empty() bool {
	return len(a.value) == 0
}
