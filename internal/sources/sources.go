package sources

import (
	"context"
	"gonstructor/internal/system"
)

type Source interface {
	LoadSystem(sys system.System)
	Listen()
}

type StateRepo interface {
	Get(ctx context.Context, initiatorID string) (string, error)
	Set(ctx context.Context, initiatorID, state string) error
}
