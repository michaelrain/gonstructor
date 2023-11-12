package domain

import (
	"context"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
)

type SourceResponder interface {
	Send(ctx context.Context, state *state.State, response response.Response) error
}
