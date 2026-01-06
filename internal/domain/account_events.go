package domain

import "yoshiyoshifujii/go-eventstore/internal/lib/eventstore"

type (
	AccountCreatedEvent struct {
		eventstore.BaseEvent
	}
	DepositedEvent struct {
		eventstore.BaseEvent
		amount uint64
	}
	WithdrawnEvent struct {
		eventstore.BaseEvent
		amount uint64
	}
	AccountClosedEvent struct {
		eventstore.BaseEvent
	}
)

func NewAccountCreatedEvent(accountID AccountID, seqNr eventstore.SeqNr) eventstore.Event {
	return AccountCreatedEvent{
		BaseEvent: eventstore.NewBaseEvent(accountID, seqNr),
	}
}

func (ev AccountCreatedEvent) EventTypeName() string {
	return "AccountCreatedEvent"
}

func (ev AccountCreatedEvent) WithSeqNr(seqNr eventstore.SeqNr) eventstore.Event {
	ev.BaseEvent = ev.BaseEvent.WithSeqNr(seqNr)
	return ev
}

func (ev AccountCreatedEvent) IsCreated() bool {
	return true
}

func (ev AccountCreatedEvent) Empty() bool {
	return false
}

func NewDepositedEvent(accountID AccountID, seqNr eventstore.SeqNr, amount uint64) eventstore.Event {
	return DepositedEvent{
		BaseEvent: eventstore.NewBaseEvent(accountID, seqNr),
		amount:    amount,
	}
}

func (ev DepositedEvent) EventTypeName() string {
	return "DepositedEvent"
}

func (ev DepositedEvent) WithSeqNr(seqNr eventstore.SeqNr) eventstore.Event {
	ev.BaseEvent = ev.BaseEvent.WithSeqNr(seqNr)
	return ev
}

func (ev DepositedEvent) IsCreated() bool {
	return false
}

func (ev DepositedEvent) Empty() bool {
	return ev.AggregateID().Empty()
}

func NewWithdrawnEvent(accountID AccountID, seqNr eventstore.SeqNr, amount uint64) WithdrawnEvent {
	return WithdrawnEvent{
		BaseEvent: eventstore.NewBaseEvent(accountID, seqNr),
		amount:    amount,
	}
}

func (ev WithdrawnEvent) EventTypeName() string {
	return "WithdrawnEvent"
}

func (ev WithdrawnEvent) WithSeqNr(seqNr eventstore.SeqNr) eventstore.Event {
	ev.BaseEvent = ev.BaseEvent.WithSeqNr(seqNr)
	return ev
}

func (ev WithdrawnEvent) IsCreated() bool {
	return false
}

func (ev WithdrawnEvent) Empty() bool {
	return ev.AggregateID().Empty()
}

func NewAccountClosedEvent(accountID AccountID, seqNr eventstore.SeqNr) AccountClosedEvent {
	return AccountClosedEvent{
		BaseEvent: eventstore.NewBaseEvent(accountID, seqNr),
	}
}

func (ev AccountClosedEvent) EventTypeName() string {
	return "AccountClosedEvent"
}

func (ev AccountClosedEvent) WithSeqNr(seqNr eventstore.SeqNr) eventstore.Event {
	ev.BaseEvent = ev.BaseEvent.WithSeqNr(seqNr)
	return ev
}

func (ev AccountClosedEvent) IsCreated() bool {
	return false
}

func (ev AccountClosedEvent) Empty() bool {
	return false
}
