package components

import (
	"context"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
	"net/http"
	"testing"
)

var testRequest = RequestComponent{
	uri:    "https://api.genderize.io/",
	method: "GET",
	vars: map[string]string{
		"%COUNT%":  "$.count",
		"%GENDER%": "$.gender",
	},
	query: map[string]string{
		"name": "michael",
	},
	client: &http.Client{},
}

var testState = state.State{
	UserID:        "test",
	Command:       "test",
	Message:       "test",
	Source:        "test",
	ModuleContext: map[string]interface{}{},
	DataBag:       map[string]interface{}{},
}

func TestA(t *testing.T) {
	ctx := context.Background()
	res := response.Response{}

	testRequest.Process(ctx, &testState, &res)

	if testState.DataBag["%GENDER%"] != "male" {
		t.Error("error")
	}

}
