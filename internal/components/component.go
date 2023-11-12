package components

import (
	"context"
	"gonstructor/internal/domain"
	"gonstructor/internal/domain/response"
)

type BaseComponent struct {
	Code          string
	ComponentType string
	next          domain.Component
}

func (component BaseComponent) Process(ctx context.Context, response *response.Response) error {
	return nil
}

func (component *BaseComponent) SetNext(c domain.Component) {
	component.next = c
}
