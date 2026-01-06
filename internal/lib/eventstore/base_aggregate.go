package eventstore

type BaseAggregate struct {
	aggregateID     AggregateID
	seqNr           SeqNr
	snapshotVersion uint64
}

func NewBaseAggregate(aggregateID AggregateID, seqNr SeqNr, snapshotVersion uint64) BaseAggregate {
	return BaseAggregate{
		aggregateID:     aggregateID,
		seqNr:           seqNr,
		snapshotVersion: snapshotVersion,
	}
}

func (a BaseAggregate) AggregateID() AggregateID {
	return a.aggregateID
}

func (a BaseAggregate) SeqNr() SeqNr {
	return a.seqNr
}

func (a BaseAggregate) SnapshotVersion() uint64 {
	return a.snapshotVersion
}

func (a BaseAggregate) WithSnapshotVersion(v uint64) BaseAggregate {
	a.snapshotVersion = v
	return a
}

func (a BaseAggregate) WithSeqNr(seqNr SeqNr) BaseAggregate {
	a.seqNr = seqNr
	return a
}
