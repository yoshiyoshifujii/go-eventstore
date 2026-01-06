package eventstore

import "context"

type (
	reader interface {
		GetLatestSnapshotByID(ctx context.Context, aggregateID AggregateID) (*AggregateResult, error)
		GetEventsByIDSinceSeqNr(ctx context.Context, aggregateID AggregateID, seqNr uint64) ([]Event, error)
	}

	writer interface {
		PersistEventAndSnapshot(ctx context.Context, event Event, aggregate Aggregate) error
	}

	EventStore interface {
		reader
		writer
	}
)
