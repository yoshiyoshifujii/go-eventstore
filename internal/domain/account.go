package domain

import "yoshiyoshifujii/go-eventstore/internal/lib/eventstore"

type (
	Account interface {
		eventstore.Aggregate
	}
)

const (
	Zero = uint64(0)
)
