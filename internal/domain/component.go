package domain

import (
	"context"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
)

type Component interface {
	Process(ctx context.Context, state *state.State, response *response.Response) error
	SetNext(component Component)
}
