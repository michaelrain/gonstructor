package scenario

import (
	"fmt"
	"gonstructor/internal/components"
	"gonstructor/internal/domain"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseActions(path string) (Actions map[string]Action, err error) {
	data := Data{}
	Actions = map[string]Action{}

	if path != "" {
		path = "./scenario.yaml"
	}

	dataByte, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(dataByte, &data)

	fmt.Println(string(dataByte))

	for _, action := range data.Actions {
		scr := Action{
			Code: action.Code,
		}

		var current domain.Component
		var next domain.Component

		for i := len(action.Components) - 1; i >= 0; i-- {
			componentData := action.Components[i]
			current = parseComponent(componentData)

			if current == nil {
				continue
			}

			if next != nil {
				current.SetNext(next)
			} else {
				trailing := components.NewTrailingComponent()
				current.SetNext(&trailing)
			}

			next = current
			scr.Components = append(scr.Components, current)
		}
		scr.FirstComponent = next
		Actions[scr.Code] = scr
	}

	return Actions, err
}

type Data struct {
	Sources map[string]interface{}
	Buttons map[string]Buttons
	Actions []struct {
		Code       string          `json:"code" yaml:"code"`
		Components []ComponentData `json:"components" yaml:"components"`
	} `json:"actions" yaml:"actions"`
}

// contains any data that a particular component may contain
type ComponentData struct {
	Code        string            `json:"code" yaml:"code"`
	MessageText string            `json:"message_text,omitempty" yaml:"message_text,omitempty"`
	Variable    string            `json:"variable,omitempty" yaml:"variable,omitempty"`
	Vars        map[string]string `yaml:"vars" json:"vars"`
	Query       map[string]string `yaml:"query" json:"query"`
	Headers     map[string]string `yaml:"headers" json:"headers"`
	URI         string            `yaml:"uri" json:"uri"`
	Method      string            `yaml:"method" json:"method"`
	Buttons     [][]Buttons       `yaml:"buttons,omitempty"`
}

type Buttons struct {
	Action string `yaml:"action" json:"action"`
	Code   string `yaml:"code" json:"code"`
	Text   string `yaml:"text" json:"text"`
}

func parseComponent(componentData ComponentData) domain.Component {
	switch componentData.Code {
	case "text":
		return parseTextComponent(componentData)
	case "buttons":
		return parseButtonsComponent(componentData)
	case "input":
		return parseInputComponent(componentData)
	case "request":
		return parseRequestComponent(componentData)
	}

	return nil
}

func parseTextComponent(componentData ComponentData) domain.Component {
	return &components.MessageComponent{
		MessageText: componentData.MessageText,
	}
}

func parseButtonsComponent(componentData ComponentData) domain.Component {
	btns := [][]components.Button{}

	for _, line := range componentData.Buttons {
		l := []components.Button{}
		for _, button := range line {
			l = append(l, components.Button{
				Code:    button.Code,
				Name:    button.Code,
				Display: button.Text,
				Action:  button.Action,
			})
		}

		btns = append(btns, l)
	}

	c := components.ButtonsComponent{
		Buttons: btns,
	}

	return &c
}

func parseInputComponent(componentData ComponentData) domain.Component {
	return &components.InputComponent{
		Code:    componentData.Code,
		Var:     componentData.Variable,
		Message: componentData.MessageText,
	}
}

func parseRequestComponent(componentData ComponentData) domain.Component {
	return components.NewRequestComponent(
		componentData.URI,
		componentData.Method,
		componentData.Headers,
		componentData.Query,
		componentData.Vars)
}
