package components

import (
	"context"
	"fmt"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
)

type ButtonsComponent struct {
	BaseComponent
	Buttons [][]Button `json:"buttons"`
}

type Button struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Display string `json:"display"`
	Action  string `json:"action"`
}

func (component ButtonsComponent) Process(ctx context.Context, s *state.State, resp *response.Response) error {

	buttons := [][]response.Button{}
	for _, l := range component.Buttons {
		line := []response.Button{}
		for _, v := range l {
			fmt.Println(v.Action)
			b := response.Button{
				Code:    v.Code,
				Display: v.Display,
				Target: response.ButtonTarget{
					Type:     "screen",
					Resource: v.Action,
				},
			}
			line = append(line, b)
		}
		buttons = append(buttons, line)
	}

	resp.Messages[0].Buttons = buttons

	return component.next.Process(ctx, s, resp)
}
