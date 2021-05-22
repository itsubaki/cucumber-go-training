package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

type apiFeature struct {
	resp *httptest.ResponseRecorder
}

func (a *apiFeature) reset(m *messages.Pickle) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) SendRequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch endpoint {
	case "/version":
		getVersion(a.resp, req)
	default:
		err = fmt.Errorf("unknown endpoint: %s", endpoint)
	}
	return
}

func (a *apiFeature) CodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) ShouldMatchJSON(arg1 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}



func FeatureContext(s *godog.Suite) {
	api := &apiFeature{}

	s.BeforeScenario(api.reset)
	s.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, api.SendRequestTo)
	s.Step(`^the response code should be (\d+)$`, api.CodeShouldBe)
	s.Step(`^the response should match json:$`, api.ShouldMatchJSON)

}
