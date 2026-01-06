package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"yoshiyoshifujii/go-eventstore/internal/domain"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

func TestOnMemoryRepository_ReplaysEvents(t *testing.T) {
	ctx := context.Background()
	accountID := domain.NewAccountID("account-1")

	store := eventstore.NewOnMemoryEventStore()
	repo := eventstore.NewRepository(store, func() eventstore.Aggregate {
		return domain.NewBlankAccount(accountID)
	})

	aggregate := eventstore.Aggregate(domain.NewBlankAccount(accountID))
	assert.Equal(t, uint64(0), aggregate.SeqNr())
	assert.Equal(t, uint64(1), aggregate.SnapshotVersion())

	var err error
	aggregate, err = repo.Save(ctx, domain.NewCreateAccountCommand(accountID), aggregate)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), aggregate.SeqNr())
	assert.Equal(t, uint64(1), aggregate.SnapshotVersion())
	aggregate, err = repo.Save(ctx, domain.NewDepositCommand(accountID, 100), aggregate)
	require.NoError(t, err)
	assert.Equal(t, uint64(2), aggregate.SeqNr())
	assert.Equal(t, uint64(1), aggregate.SnapshotVersion())
	aggregate, err = repo.Save(ctx, domain.NewWithdrawCommand(accountID, 100), aggregate)
	require.NoError(t, err)
	assert.Equal(t, uint64(3), aggregate.SeqNr())
	assert.Equal(t, uint64(1), aggregate.SnapshotVersion())
	aggregate, err = repo.Save(ctx, domain.NewCloseAccountCommand(accountID), aggregate)
	require.NoError(t, err)
	assert.Equal(t, uint64(4), aggregate.SeqNr())
	assert.Equal(t, uint64(1), aggregate.SnapshotVersion())

	loaded, err := repo.FindBy(ctx, accountID)
	require.NoError(t, err)
	require.NotNil(t, loaded)
	require.Equal(t, "closed_account", (*loaded).AggregateTypeName())
	require.Equal(t, uint64(4), (*loaded).SeqNr())
	require.Equal(t, uint64(1), (*loaded).SnapshotVersion())
}
