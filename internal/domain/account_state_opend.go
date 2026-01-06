package domain

import (
	"errors"
	"fmt"
	"yoshiyoshifujii/go-eventstore/internal/lib/eventstore"
)

type (
	OpenedAccount struct {
		id              AccountID
		seqNr           uint64
		snapshotVersion uint64
		balance         uint64
	}
)

func NewOpenedAccount(id AccountID, seqNr, snapshotVersion, balance uint64) OpenedAccount {
	return OpenedAccount{
		id:              id,
		seqNr:           seqNr,
		snapshotVersion: snapshotVersion,
		balance:         balance,
	}
}

func (ag OpenedAccount) AggregateID() eventstore.AggregateID {
	return ag.id
}

func (ag OpenedAccount) AggregateTypeName() string {
	return "opened_account"
}

func (ag OpenedAccount) SeqNr() uint64 {
	return ag.seqNr
}

func (ag OpenedAccount) SnapshotVersion() uint64 {
	return ag.snapshotVersion
}

func (ag OpenedAccount) WithVersion(v uint64) eventstore.Aggregate {
	return ag.WithSnapshotVersion(v)
}

func (ag OpenedAccount) WithSnapshotVersion(v uint64) eventstore.Aggregate {
	ag.snapshotVersion = v
	return ag
}

func (ag OpenedAccount) ApplyCommand(command eventstore.Command) (eventstore.Event, error) {
	switch cmd := command.(type) {
	case DepositCommand:
		return NewDepositedEvent(ag.id, ag.seqNr+1, cmd.amount), nil
	case WithdrawCommand:
		if ag.canWithdraw(cmd.amount) {
			return NewWithdrawnEvent(ag.id, ag.seqNr+1, cmd.amount), nil
		}
		return nil, fmt.Errorf("insufficient balance %d to be able to withdraw %d", ag.balance, cmd.amount)
	case CloseAccountCommand:
		if ag.balance == Zero {
			return NewAccountClosedEvent(ag.id, ag.seqNr+1), nil
		}
		return nil, errors.New("can't close account with non-zero balance")
	default:
		return nil, errors.New("unknown command")
	}
}

func (ag OpenedAccount) ApplyEvent(event eventstore.Event) eventstore.Aggregate {
	switch ev := event.(type) {
	case AccountCreatedEvent:
		panic("unsupported event")
	case DepositedEvent:
		ag.balance = ag.balance + ev.amount
		return ag
	case WithdrawnEvent:
		ag.balance = ag.balance - ev.amount
		return ag
	case AccountClosedEvent:
		return NewClosedAccount(ev.accountID, ag.seqNr)
	default:
		panic("unknown event")
	}
}

func (ag OpenedAccount) Empty() bool {
	return ag.id.Empty()
}

func (ag OpenedAccount) canWithdraw(amount uint64) bool {
	return ag.balance-amount >= Zero
}
