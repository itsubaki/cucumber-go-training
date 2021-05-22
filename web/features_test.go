package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/cucumber/godog"
	"github.com/jfilipczyk/gomatch"
)

var api = &apiFeature{}

type apiFeature struct {
	header http.Header
	body   io.Reader
	resp   *httptest.ResponseRecorder
	keep   map[string]interface{}
}

func (a *apiFeature) start() {
	a.keep = make(map[string]interface{})
}

func (a *apiFeature) stop() {
}

func (a *apiFeature) reset(sc *godog.Scenario) {
	a.header = make(http.Header)
	a.body = nil
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) replace(str string) string {
	for k, v := range a.keep {
		switch val := v.(type) {
		case string:
			str = strings.Replace(str, k, val, -1)
		default:
			continue
		}
	}

	return str
}

func (a *apiFeature) SetHeader(k, v string) error {
	a.header.Add(k, v)
	return nil
}

func (a *apiFeature) SetRequestBody(body *godog.DocString) error {
	r := a.replace(body.Content)
	a.body = bytes.NewBuffer([]byte(r))
	return nil
}

func (a *apiFeature) Request(method, endpoint string) error {
	r := a.replace(endpoint)
	req := httptest.NewRequest(method, r, a.body)
	req.Header = a.header

	a.resp.Code = 200
	return nil
}

func (a *apiFeature) ResponseCodeShouldBe(code int) error {
	if code == a.resp.Code {
		return nil
	}

	return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
}

func (a *apiFeature) ResponseShouldMatchJSON(body *godog.DocString) error {
	expected := a.replace(body.Content)
	actual := a.resp.Body.String()

	ok, err := gomatch.NewDefaultJSONMatcher().Match(expected, actual)
	if err != nil {
		return fmt.Errorf("actual=%s, match: %v", actual, err)
	}

	if !ok {
		return fmt.Errorf("expected JSON does not match actual, %s vs. %s", expected, actual)
	}

	return nil
}

func (a *apiFeature) Keep(key, as string) error {
	var actual map[string]interface{}
	if err := json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return fmt.Errorf("body=%s, unmarshal: %v", a.resp.Body.String(), err)
	}

	if v, ok := actual[key]; ok {
		a.keep[as] = v
	}

	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	before := func() {
		api.start()
	}
	after := func() {
		api.stop()
	}

	ctx.BeforeSuite(before)
	ctx.AfterSuite(after)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(api.reset)

	ctx.Step(`^I set "([^"]*)" header with "([^"]*)"$`, api.SetHeader)
	ctx.Step(`^I set request body:$`, api.SetRequestBody)
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.Request)
	ctx.Step(`^the response code should be (\d+)$`, api.ResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, api.ResponseShouldMatchJSON)
	ctx.Step(`^I keep the JSON response at "([^"]*)" as "([^"]*)"$`, api.Keep)
}
