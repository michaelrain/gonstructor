package system

import (
	"context"
	"fmt"
	"gonstructor/internal/components"
	"gonstructor/internal/domain"
	"gonstructor/internal/domain/state"
	"gonstructor/internal/scenario"
	"gonstructor/internal/triggers"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

type System struct {
	StateRepository domain.StateRepository
	actions         map[string]scenario.Action
	triggers        map[string]triggers.Trigger
	responders      map[string]domain.SourceResponder
	chanRequest     chan domain.Request
	logger          *logrus.Logger
}

const ModuleNotFoundError = "module not found"

func NewSystem(stateRepo domain.StateRepository, actions map[string]scenario.Action, responders map[string]domain.SourceResponder, chanRequest chan domain.Request, logger *logrus.Logger) (*System, error) {
	sys := System{
		logger:          logger,
		actions:         actions,
		responders:      responders,
		chanRequest:     chanRequest,
		StateRepository: stateRepo,
	}
	sys.InitTriggers()

	return &sys, nil
}

func (sys *System) InitTriggers() {
	var t map[string]triggers.Trigger
	for _, screen := range sys.actions {
		for _, v := range screen.Components {
			if reflect.TypeOf(v).String() == "components.ButtonsComponent" {
				bc := v.(*components.ButtonsComponent)
				for _, line := range bc.Buttons {
					for _, button := range line {
						triggers.RegisterTrigger(button.Code, "button")
					}
				}
			}
		}
	}
	sys.triggers = t
}

func (sys *System) Listen(ctx context.Context) {
	for {
		select {
		case request := <-sys.chanRequest:
			sys.logger.Trace("handle request")
			s, err := sys.GetCurrentState(ctx, fmt.Sprintf("%s:%s", request.Source, request.FromID))

			if err != nil {
				s = getBaseStateFromRequest(request)
			}

			command := request.Command
			if s.CapturedBy != "" {
				c := strings.Split(s.CapturedBy, ".")
				command = c[0]
			}

			responder, ok := sys.responders[s.Source]

			if !ok {
				sys.logger.Error("not found responder " + s.Source)
				return
			}

			screen, ok := sys.actions[command]

			if !ok {
				sys.logger.Error("comand not found " + command)
				return
			}

			err = screen.Render(ctx, &s, responder)

			if err != nil {
				fmt.Println(err.Error())
				// @todo log
				return
			}

			sys.StateRepository.Set(ctx, request.FromID, s)
		case <-ctx.Done():
			return
		}
	}
}

func getBaseStateFromRequest(req domain.Request) state.State {
	return state.State{
		UserID:  req.FromID,
		Source:  req.Source,
		Message: req.Command,
		Command: req.Command,
	}
}

func (sys *System) GetCurrentState(ctx context.Context, initiatorID string) (state.State, error) {
	return sys.StateRepository.Get(ctx, initiatorID)
}
