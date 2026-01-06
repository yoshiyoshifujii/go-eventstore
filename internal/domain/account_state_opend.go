package domain

import (
	"errors"
	"fmt"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

type (
	OpenedAccount struct {
		eventstore.BaseAggregate
		balance uint64
	}
)

func NewOpenedAccount(id AccountID, seqNr eventstore.SeqNr, snapshotVersion, balance uint64) OpenedAccount {
	return OpenedAccount{
		BaseAggregate: eventstore.NewBaseAggregate(id, seqNr, snapshotVersion),
		balance:       balance,
	}
}

func (ag OpenedAccount) AggregateTypeName() string {
	return "opened_account"
}

func (ag OpenedAccount) WithSnapshotVersion(v uint64) eventstore.Aggregate {
	ag.BaseAggregate = ag.BaseAggregate.WithSnapshotVersion(v)
	return ag
}

func (ag OpenedAccount) ApplyCommand(command eventstore.Command) (eventstore.Event, error) {
	switch cmd := command.(type) {
	case DepositCommand:
		return NewDepositedEvent(ag.AggregateID().(AccountID), ag.SeqNr(), cmd.amount), nil
	case WithdrawCommand:
		if ag.canWithdraw(cmd.amount) {
			return NewWithdrawnEvent(ag.AggregateID().(AccountID), ag.SeqNr(), cmd.amount), nil
		}
		return nil, fmt.Errorf("insufficient balance %d to be able to withdraw %d", ag.balance, cmd.amount)
	case CloseAccountCommand:
		if ag.balance == Zero {
			return NewAccountClosedEvent(ag.AggregateID().(AccountID), ag.SeqNr()), nil
		}
		return nil, errors.New("can't close account with non-zero balance")
	default:
		return nil, errors.New("unknown command")
	}
}

func (ag OpenedAccount) ApplyEvent(event eventstore.Event) eventstore.Aggregate {
	switch ev := event.(type) {
	case AccountCreatedEvent:
		panic("unsupported event")
	case DepositedEvent:
		ag.BaseAggregate = ag.BaseAggregate.WithSeqNr(ev.SeqNr())
		ag.balance = ag.balance + ev.amount
		return ag
	case WithdrawnEvent:
		ag.BaseAggregate = ag.BaseAggregate.WithSeqNr(ev.SeqNr())
		ag.balance = ag.balance - ev.amount
		return ag
	case AccountClosedEvent:
		return NewClosedAccount(ev.AggregateID().(AccountID), ev.SeqNr(), ag.SnapshotVersion())
	default:
		panic("unknown event")
	}
}

func (ag OpenedAccount) Empty() bool {
	return ag.AggregateID().Empty()
}

func (ag OpenedAccount) canWithdraw(amount uint64) bool {
	return ag.balance-amount >= Zero
}
