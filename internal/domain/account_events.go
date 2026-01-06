package domain

import (
	"fmt"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

type (
	AccountCreatedEvent struct {
		accountID AccountID
		seqNr     uint64
	}
	DepositedEvent struct {
		accountID AccountID
		seqNr     uint64
		amount    uint64
	}
	WithdrawnEvent struct {
		accountID AccountID
		seqNr     uint64
		amount    uint64
	}
	AccountClosedEvent struct {
		accountID AccountID
		seqNr     uint64
	}
)

func generateID(ev eventstore.Event) string {
	return fmt.Sprintf("%s-%d", ev.AggregateID().Value(), ev.SeqNr())
}

func NewAccountCreatedEvent(accountID AccountID, seqNr uint64) eventstore.Event {
	return AccountCreatedEvent{
		accountID: accountID,
		seqNr:     seqNr,
	}
}

func (ev AccountCreatedEvent) ID() string {
	return generateID(ev)
}

func (ev AccountCreatedEvent) EventTypeName() string {
	return "AccountCreatedEvent"
}

func (ev AccountCreatedEvent) AggregateID() eventstore.AggregateID {
	return ev.accountID
}

func (ev AccountCreatedEvent) SeqNr() uint64 {
	return ev.seqNr
}

func (ev AccountCreatedEvent) IsCreated() bool {
	return true
}

func (ev AccountCreatedEvent) Empty() bool {
	return false
}

func NewDepositedEvent(accountID AccountID, seqNr uint64, amount uint64) eventstore.Event {
	return DepositedEvent{
		accountID: accountID,
		seqNr:     seqNr,
		amount:    amount,
	}
}

func (ev DepositedEvent) ID() string {
	return generateID(ev)
}

func (ev DepositedEvent) EventTypeName() string {
	return "DepositedEvent"
}

func (ev DepositedEvent) AggregateID() eventstore.AggregateID {
	return ev.accountID
}

func (ev DepositedEvent) SeqNr() uint64 {
	return ev.seqNr
}

func (ev DepositedEvent) IsCreated() bool {
	return false
}

func (ev DepositedEvent) Empty() bool {
	return ev.AggregateID().Empty()
}

func NewWithdrawnEvent(accountID AccountID, seqNr uint64, amount uint64) WithdrawnEvent {
	return WithdrawnEvent{
		accountID: accountID,
		seqNr:     seqNr,
		amount:    amount,
	}
}

func (ev WithdrawnEvent) ID() string {
	return generateID(ev)
}

func (ev WithdrawnEvent) EventTypeName() string {
	return "WithdrawnEvent"
}

func (ev WithdrawnEvent) AggregateID() eventstore.AggregateID {
	return ev.accountID
}

func (ev WithdrawnEvent) SeqNr() uint64 {
	return ev.seqNr
}

func (ev WithdrawnEvent) IsCreated() bool {
	return false
}

func (ev WithdrawnEvent) Empty() bool {
	return ev.AggregateID().Empty()
}

func NewAccountClosedEvent(accountID AccountID, seqNr uint64) AccountClosedEvent {
	return AccountClosedEvent{
		accountID: accountID,
		seqNr:     seqNr,
	}
}

func (ev AccountClosedEvent) ID() string {
	return generateID(ev)
}

func (ev AccountClosedEvent) EventTypeName() string {
	return "AccountClosedEvent"
}

func (ev AccountClosedEvent) AggregateID() eventstore.AggregateID {
	return ev.accountID
}

func (ev AccountClosedEvent) SeqNr() uint64 {
	return ev.seqNr
}

func (ev AccountClosedEvent) IsCreated() bool {
	return false
}

func (ev AccountClosedEvent) Empty() bool {
	return false
}
