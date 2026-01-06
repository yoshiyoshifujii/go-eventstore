package domain

import (
	"errors"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

type (
	BlankAccount struct {
		id AccountID
	}
)

func NewBlankAccount(id AccountID) BlankAccount {
	return BlankAccount{
		id: id,
	}
}

func (ag BlankAccount) AggregateID() eventstore.AggregateID {
	return ag.id
}

func (ag BlankAccount) AggregateTypeName() string {
	return "blank_account"
}

func (ag BlankAccount) SeqNr() eventstore.SeqNr {
	return eventstore.NewSeqNr(0)
}

func (ag BlankAccount) SnapshotVersion() uint64 {
	return 1
}

func (ag BlankAccount) WithVersion(v uint64) eventstore.Aggregate {
	return ag.WithSnapshotVersion(v)
}

func (ag BlankAccount) WithSnapshotVersion(v uint64) eventstore.Aggregate {
	panic("with snapshot version not supported")
}

func (ag BlankAccount) ApplyCommand(command eventstore.Command) (eventstore.Event, error) {
	switch command.(type) {
	case CreateAccountCommand:
		return NewAccountCreatedEvent(ag.id, ag.SeqNr().Next()), nil
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
	return ag.id.Empty()
}
