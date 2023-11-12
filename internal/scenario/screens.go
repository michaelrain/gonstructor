package scenario

import (
	"context"
	"fmt"
	"gonstructor/internal/domain"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
)

type Action struct {
	Triggers       []string `json:"triggers"`
	Code           string
	Components     []domain.Component
	FirstComponent domain.Component
	Middlewares    []string `json:"middlewares"`
	IsDeleted      bool     `json:"is_deleted"`
}

func (a Action) Render(ctx context.Context, st *state.State, responder domain.SourceResponder) error {
	var res response.Response
	res.Messages = []response.Message{}
	res.Messages = append(res.Messages, response.Message{})

	err := a.FirstComponent.Process(ctx, st, &res)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(responder)

	return responder.Send(ctx, st, res)
}

func GetActions() (actions map[string]Action) {

	actions, err := ParseActions("")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return actions
}
