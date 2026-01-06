package eventstore

import (
	"context"
	"sync"
)

type OnMemoryEventStore struct {
	mu        sync.RWMutex
	events    map[string][]Event
	snapshots map[string]AggregateResult
}

func NewOnMemoryEventStore() *OnMemoryEventStore {
	return &OnMemoryEventStore{
		events:    make(map[string][]Event),
		snapshots: make(map[string]AggregateResult),
	}
}

func (store *OnMemoryEventStore) GetLatestSnapshotByID(_ context.Context, aggregateID AggregateID) (*AggregateResult, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	key := aggregateID.AsString()
	result, ok := store.snapshots[key]
	if !ok {
		return &AggregateResult{}, nil
	}
	return &result, nil
}

func (store *OnMemoryEventStore) GetEventsByIDSinceSeqNr(_ context.Context, aggregateID AggregateID, seqNr uint64) ([]Event, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	key := aggregateID.AsString()
	events := store.events[key]
	if len(events) == 0 {
		return nil, nil
	}
	var filtered []Event
	for _, event := range events {
		if event.SeqNr() >= seqNr {
			filtered = append(filtered, event)
		}
	}
	return filtered, nil
}

func (store *OnMemoryEventStore) PersistEventAndSnapshot(_ context.Context, event Event, aggregate Aggregate) error {
	store.mu.Lock()
	defer store.mu.Unlock()
	key := event.AggregateID().AsString()
	store.events[key] = append(store.events[key], event)
	store.snapshots[key] = AggregateResult{Aggregate: aggregate}
	return nil
}
