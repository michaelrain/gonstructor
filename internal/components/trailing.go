package components

import (
	"context"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
)

type TrailingComponent struct {
	BaseComponent
}

func NewTrailingComponent() TrailingComponent {
	return TrailingComponent{}
}

func (component TrailingComponent) Process(ctx context.Context, s *state.State, resp *response.Response) error {
	return nil
}
