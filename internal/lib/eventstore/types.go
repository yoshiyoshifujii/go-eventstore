package eventstore

type (
	SeqNr struct {
		value uint64
	}

	AggregateID interface {
		AggregateIDTypeName() string
		Value() string
		AsString() string
		Empty() bool
	}

	Command interface {
		CommandTypeName() string
		Empty() bool
	}

	Event interface {
		ID() string
		EventTypeName() string
		AggregateID() AggregateID
		SeqNr() SeqNr
		WithSeqNr(SeqNr) Event
		IsCreated() bool
		Empty() bool
	}

	Aggregate interface {
		AggregateID() AggregateID
		AggregateTypeName() string
		SeqNr() SeqNr
		SnapshotVersion() uint64
		WithSnapshotVersion(uint64) Aggregate
		ApplyCommand(Command) (Event, error)
		ApplyEvent(Event) Aggregate
		Empty() bool
	}

	AggregateResult struct {
		Aggregate Aggregate
	}
)

func NewSeqNr(value uint64) SeqNr {
	return SeqNr{value: value}
}

func (s SeqNr) Value() uint64 {
	return s.value
}

func (s SeqNr) next() SeqNr {
	return SeqNr{value: s.value + 1}
}

func (a AggregateResult) Empty() bool {
	return a.Aggregate == nil
}
