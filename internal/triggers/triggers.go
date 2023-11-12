package triggers

import "gonstructor/internal/scenario"

type Trigger struct {
	Code    string          `json:"code"`
	Type    string          `json:"type"`
	Handler func()          `json:"handler"`
	Screen  scenario.Action `json:"action"`
}

var triggers map[string]Trigger

func init() {
	triggers = map[string]Trigger{}
}

func RegisterTrigger(code string, ty string) Trigger {
	trigger := Trigger{
		Code: code,
		Type: ty,
	}

	triggers[code] = trigger

	return trigger
}
