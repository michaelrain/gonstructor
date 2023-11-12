package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"gonstructor/internal/domain/state"

	"github.com/go-redis/redis/v8"
)

type StateRepository struct {
	client *redis.Client
}

func NewStateRepository(client *redis.Client) *StateRepository {
	return &StateRepository{
		client: client,
	}
}

func (s *StateRepository) Get(ctx context.Context, key string) (state.State, error) {
	var st state.State

	res := s.client.Get(ctx, fmt.Sprintf("context:%s", key))

	result, err := res.Result()
	fmt.Println("redis get result")
	fmt.Println(result)

	if err == redis.Nil {
		return st, nil
	}

	if err != nil {
		return st, err
	}

	err = json.Unmarshal([]byte(result), &st)

	return st, err
}

func (s *StateRepository) Set(ctx context.Context, key string, state state.State) error {
	dur, err := time.ParseDuration("12h")
	if err != nil {
		return err
	}

	stateByte, err := json.Marshal(state)

	if err != nil {
		return err
	}

	fmt.Println("redis set stateString")
	fmt.Println(string(stateByte))

	res := s.client.Set(ctx, fmt.Sprintf("context:%s", key), string(stateByte), dur)

	if _, err := res.Result(); err != nil {
		return err
	}

	return nil
}

func (s *StateRepository) Reset(ctx context.Context, key string) error {
	res := s.client.Set(ctx, fmt.Sprintf("context:%s", key), nil, 0)

	if _, err := res.Result(); err != nil {
		return err
	}

	return nil
}
