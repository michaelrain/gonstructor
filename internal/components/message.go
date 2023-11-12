package components

import (
	"context"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
	"gonstructor/internal/helpers"
)

type MessageComponent struct {
	BaseComponent
	MessageText string `json:"messageText"`
}

func (component MessageComponent) Process(ctx context.Context, s *state.State, resp *response.Response) error {
	component.MessageText = helpers.ApplyVar(component.MessageText, s.DataBag)
	resp.Messages[0].Text = component.MessageText

	component.next.Process(ctx, s, resp)

	return nil
}
