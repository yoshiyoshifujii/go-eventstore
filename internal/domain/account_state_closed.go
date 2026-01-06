package domain

import (
	"errors"
	"fmt"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

type (
	ClosedAccount struct {
		eventstore.BaseAggregate
	}
)

func NewClosedAccount(id AccountID, seqNr eventstore.SeqNr, snapshotVersion uint64) ClosedAccount {
	return ClosedAccount{
		BaseAggregate: eventstore.NewBaseAggregate(id, seqNr, snapshotVersion),
	}
}

func (ag ClosedAccount) AggregateTypeName() string {
	return "closed_account"
}

func (ag ClosedAccount) WithSnapshotVersion(v uint64) eventstore.Aggregate {
	ag.BaseAggregate = ag.BaseAggregate.WithSnapshotVersion(v)
	return ag
}

func (ag ClosedAccount) ApplyCommand(command eventstore.Command) (eventstore.Event, error) {
	return nil, errors.New("account is closed")
}

func (ag ClosedAccount) ApplyEvent(event eventstore.Event) eventstore.Aggregate {
	panic(fmt.Sprintf("unexpected event %s in state [ClosedAccount]", event.EventTypeName()))
}

func (ag ClosedAccount) Empty() bool {
	return false
}
