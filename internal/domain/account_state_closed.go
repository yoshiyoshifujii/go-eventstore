package domain

import (
	"errors"
	"fmt"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

type (
	ClosedAccount struct {
		id              AccountID
		seqNr           eventstore.SeqNr
		snapshotVersion uint64
	}
)

func NewClosedAccount(id AccountID, seqNr eventstore.SeqNr, snapshotVersion uint64) ClosedAccount {
	return ClosedAccount{
		id:              id,
		seqNr:           seqNr,
		snapshotVersion: snapshotVersion,
	}
}

func (ag ClosedAccount) AggregateID() eventstore.AggregateID {
	return ag.id
}

func (ag ClosedAccount) AggregateTypeName() string {
	return "closed_account"
}

func (ag ClosedAccount) SeqNr() eventstore.SeqNr {
	return ag.seqNr
}

func (ag ClosedAccount) SnapshotVersion() uint64 {
	return ag.snapshotVersion
}

func (ag ClosedAccount) WithVersion(v uint64) eventstore.Aggregate {
	return ag.WithSnapshotVersion(v)
}

func (ag ClosedAccount) WithSnapshotVersion(v uint64) eventstore.Aggregate {
	ag.snapshotVersion = v
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
