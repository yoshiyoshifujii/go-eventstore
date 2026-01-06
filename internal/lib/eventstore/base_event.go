package eventstore

import "fmt"

type BaseEvent struct {
	aggregateID AggregateID
	seqNr       SeqNr
}

func NewBaseEvent(aggregateID AggregateID, seqNr SeqNr) BaseEvent {
	return BaseEvent{
		aggregateID: aggregateID,
		seqNr:       seqNr,
	}
}

func (e BaseEvent) ID() string {
	return fmt.Sprintf("%s-%d", e.aggregateID.Value(), e.seqNr.Value())
}

func (e BaseEvent) AggregateID() AggregateID {
	return e.aggregateID
}

func (e BaseEvent) SeqNr() SeqNr {
	return e.seqNr
}

func (e BaseEvent) WithSeqNr(seqNr SeqNr) BaseEvent {
	e.seqNr = seqNr
	return e
}
