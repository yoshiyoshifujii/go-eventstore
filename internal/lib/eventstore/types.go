package eventstore

type (
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
		SeqNr() uint64
		IsCreated() bool
		Empty() bool
	}

	Aggregate interface {
		AggregateID() AggregateID
		AggregateTypeName() string
		SeqNr() uint64
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

func (a AggregateResult) Empty() bool {
	return a.Aggregate == nil
}
