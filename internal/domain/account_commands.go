package domain

import "yoshiyoshifujii/go-eventstore/internal/lib/eventstore"

type (
	CreateAccountCommand struct {
		accountID AccountID
	}
	DepositCommand struct {
		accountID AccountID
		amount    uint64
	}
	WithdrawCommand struct {
		accountID AccountID
		amount    uint64
	}
	CloseAccountCommand struct {
		accountID AccountID
	}
)

func NewCreateAccountCommand(accountID AccountID) CreateAccountCommand {
	return CreateAccountCommand{accountID: accountID}
}

func (c CreateAccountCommand) CommandTypeName() string {
	return "CreateAccount"
}

func (c CreateAccountCommand) AggregateID() eventstore.AggregateID {
	return c.accountID
}

func (c CreateAccountCommand) Empty() bool {
	return false
}

func NewDepositCommand(accountID AccountID, amount uint64) DepositCommand {
	return DepositCommand{accountID: accountID, amount: amount}
}

func (c DepositCommand) CommandTypeName() string {
	return "Deposit"
}

func (c DepositCommand) AggregateID() eventstore.AggregateID {
	return c.accountID
}

func (c DepositCommand) Empty() bool {
	return false
}

func NewWithdrawCommand(accountID AccountID, amount uint64) WithdrawCommand {
	return WithdrawCommand{accountID: accountID, amount: amount}
}

func (c WithdrawCommand) CommandTypeName() string {
	return "Withdraw"
}

func (c WithdrawCommand) AggregateID() eventstore.AggregateID {
	return c.accountID
}

func (c WithdrawCommand) Empty() bool {
	return false
}

func NewCloseAccountCommand(accountID AccountID) CloseAccountCommand {
	return CloseAccountCommand{accountID: accountID}
}

func (c CloseAccountCommand) CommandTypeName() string {
	return "CloseAccount"
}

func (c CloseAccountCommand) AggregateID() eventstore.AggregateID {
	return c.accountID
}

func (c CloseAccountCommand) Empty() bool {
	return false
}
