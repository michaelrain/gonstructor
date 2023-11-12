package domain

import (
	"context"
	"gonstructor/internal/domain/state"
)

type StateRepository interface {
	Get(ctx context.Context, key string) (state.State, error)
	Set(ctx context.Context, key string, state state.State) error
}
