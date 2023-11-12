package components

import (
	"context"
	"fmt"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
)

type InputComponent struct {
	BaseComponent

	Code    string
	Var     string
	Message string
}

func (component InputComponent) Process(ctx context.Context, s *state.State, resp *response.Response) error {
	captureKey := fmt.Sprintf("%s.%s", s.Command, component.Code)

	if s.CapturedBy == captureKey {
		if s.DataBag == nil {
			s.DataBag = map[string]interface{}{}
		}

		fmt.Println(component.Var)
		fmt.Println(s.Message)

		s.DataBag[component.Var] = s.Message
		s.CapturedBy = ""
		return component.next.Process(ctx, s, resp)
	}

	s.CapturedBy = captureKey
	fmt.Println(captureKey)
	resp.Messages[0].Text = component.Message

	return nil
}
