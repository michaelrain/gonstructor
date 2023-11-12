package components

import (
	"context"
	"encoding/json"
	"fmt"
	"gonstructor/internal/domain/response"
	"gonstructor/internal/domain/state"
	"gonstructor/internal/helpers"
	"io"
	"net/http"

	"github.com/yalp/jsonpath"
)

type RequestComponent struct {
	BaseComponent

	uri    string
	method string

	client  *http.Client
	headers map[string]string
	query   map[string]string
	vars    map[string]string
}

func NewRequestComponent(uri string, method string, headers map[string]string, query map[string]string, vars map[string]string) *RequestComponent {
	client := http.Client{}

	return &RequestComponent{
		uri:     uri,
		method:  method,
		client:  &client,
		headers: headers,
		query:   query,
		vars:    vars,
	}
}

func (component RequestComponent) Process(ctx context.Context, s *state.State, resp *response.Response) error {
	// replace %VAR% in header and query params
	component.headers = helpers.ApplyMap(component.headers, s.DataBag)
	component.query = helpers.ApplyMap(component.query, s.DataBag)

	req, err := http.NewRequest(component.method, component.uri, nil)

	if err != nil {
		return err
	}

	q := req.URL.Query()

	for k, v := range component.query {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	for k, v := range component.headers {
		req.Header.Add(k, v)
	}

	res, err := component.client.Do(req)

	if err != nil {
		return err
	}

	resByte, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var resInterface interface{}
	err = json.Unmarshal(resByte, &resInterface)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for variable, query := range component.vars {
		val, err := jsonpath.Read(resInterface, query)

		if val == nil {
			continue
		}

		if err != nil {
			continue
		}

		s.DataBag[variable] = val
	}

	return component.next.Process(ctx, s, resp)
}
