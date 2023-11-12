package domain

import (
	"context"
	"gonstructor/internal/domain/state"
)

type System interface {
	ExecuteCommand(ctx context.Context, responder SourceResponder, s state.State, command string)
}
