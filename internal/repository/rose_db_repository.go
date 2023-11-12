package repository

import (
	"context"
	"encoding/json"
	"gonstructor/internal/domain/state"

	"github.com/rosedblabs/rosedb/v2"
)

type StateRepositoryRose struct {
	db *rosedb.DB
}

func NewRoseDBStateRepository(db *rosedb.DB) StateRepositoryRose {
	return StateRepositoryRose{
		db: db,
	}
}

func (s StateRepositoryRose) Get(ctx context.Context, key string) (state.State, error) {
	state := state.State{}
	respBytes, err := s.db.Get([]byte(key))

	if err != nil {
		return state, err
	}

	err = json.Unmarshal(respBytes, &state)

	if err != nil {
		return state, err
	}

	return state, err
}

func (s StateRepositoryRose) Set(ctx context.Context, key string, state state.State) error {
	stateByte, err := json.Marshal(state)

	if err != nil {
		return err
	}

	return s.db.Put([]byte(key), stateByte)
}
