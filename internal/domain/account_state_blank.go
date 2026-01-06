package domain

import (
	"errors"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

type (
	BlankAccount struct {
		eventstore.BaseAggregate
	}
)

func NewBlankAccount(id AccountID) BlankAccount {
	return BlankAccount{
		BaseAggregate: eventstore.NewBaseAggregate(id, eventstore.NewSeqNr(0), 1),
	}
}

func (ag BlankAccount) AggregateTypeName() string {
	return "blank_account"
}

func (ag BlankAccount) WithSnapshotVersion(v uint64) eventstore.Aggregate {
	ag.BaseAggregate = ag.BaseAggregate.WithSnapshotVersion(v)
	return ag
}

func (ag BlankAccount) ApplyCommand(command eventstore.Command) (eventstore.Event, error) {
	switch command.(type) {
	case CreateAccountCommand:
		return NewAccountCreatedEvent(ag.AggregateID().(AccountID), ag.SeqNr()), nil
	default:
		return nil, errors.New("unknown command type")
	}
}

func (ag BlankAccount) ApplyEvent(event eventstore.Event) eventstore.Aggregate {
	switch ev := event.(type) {
	case AccountCreatedEvent:
		return NewOpenedAccount(ev.AggregateID().(AccountID), ev.SeqNr(), ag.SnapshotVersion(), Zero)
	default:
		panic("unknown event type")
	}
}

func (ag BlankAccount) Empty() bool {
	return ag.AggregateID().Empty()
}
