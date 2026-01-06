package eventstore

import "context"

type (
	Repository struct {
		es          EventStore
		createBlank func() Aggregate
	}
)

func NewRepository(es EventStore, createBlank func() Aggregate) Repository {
	return Repository{
		es:          es,
		createBlank: createBlank,
	}
}

func (repo Repository) FindBy(ctx context.Context, aggregateID AggregateID) (Aggregate, error) {
	if repo.es == nil {
		panic("no event store found")
	}
	if aggregateID == nil || aggregateID.Empty() {
		panic("aggregateID is required")
	}
	result, err := repo.es.GetLatestSnapshotByID(ctx, aggregateID)
	if err != nil {
		return nil, err
	}

	ag, seqNr := repo.getOrCreate(result)

	events, err := repo.es.GetEventsByIDSinceSeqNr(ctx, aggregateID, seqNr)
	if err != nil {
		return nil, err
	}

	aggregate := repo.eventHandler(ag, events)
	return aggregate, nil
}

func (repo Repository) getOrCreate(result *AggregateResult) (Aggregate, uint64) {
	var (
		ag    Aggregate
		seqNr uint64
	)
	if result.Empty() {
		ag = repo.createBlank()
		seqNr = 0
	} else {
		ag = result.Aggregate
		seqNr = result.Aggregate.SeqNr() + 1
	}
	return ag, seqNr
}

func (repo Repository) eventHandler(aggregate Aggregate, events []Event) Aggregate {
	result := aggregate
	for _, event := range events {
		result = result.ApplyEvent(event)
	}
	return result
}

func (repo Repository) Save(ctx context.Context, command Command, aggregate Aggregate) (Aggregate, error) {
	if repo.es == nil {
		panic("no event store found")
	}
	if command == nil || command.Empty() {
		panic("command is required")
	}
	if aggregate == nil || aggregate.Empty() {
		panic("aggregate is required")
	}
	event, err := aggregate.ApplyCommand(command)
	if err != nil {
		return nil, err
	}
	err = repo.es.PersistEventAndSnapshot(ctx, event, aggregate)
	if err != nil {
		return nil, err
	}
	nexAggregate := aggregate.ApplyEvent(event)
	return nexAggregate, nil
}
