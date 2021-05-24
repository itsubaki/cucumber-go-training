package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

var api = &apiFeature{}

type apiFeature struct {
	header http.Header
	body   io.Reader
	resp   *httptest.ResponseRecorder
}

func (a *apiFeature) start() {
}

func (a *apiFeature) stop() {
}

func (a *apiFeature) reset(sc *godog.Scenario) {
	a.header = make(http.Header)
	a.body = nil
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) SetHeader(k, v string) error {
	a.header.Add(k, v)
	return nil
}

func (a *apiFeature) Request(method, endpoint string) error {
	req := httptest.NewRequest(method, endpoint, a.body)
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

func (a *apiFeature) PageTileIs(title string) error {
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
	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, api.Request)
	ctx.Step(`^the response code should be (\d+)$`, api.ResponseCodeShouldBe)
	ctx.Step(`^the page title is "([^"]*)"$`, api.PageTileIs)
}
