package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cucumber/godog"
)

type scenarioData struct {
	productUrl string
	statusCode int
}

func (s *scenarioData) iSendRequestToWithAboveProductUrlInBody(method, url string) error {
	req, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(fmt.Sprintf(`{"url":"%s"}`, s.productUrl)))
	if err != nil {
		return err
	}
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	s.statusCode = response.StatusCode
	return nil
}
func (s *scenarioData) theProductUrl(url string) error {
	s.productUrl = url
	return nil
}

func (s *scenarioData) theResponseCodeShouldBe(expectedResponseCode int) error {
	if s.statusCode != expectedResponseCode {
		return fmt.Errorf("%d is not equal to %d", s.statusCode, expectedResponseCode)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	var s scenarioData
	ctx.Step(`^i send "([^"]*)" request to "([^"]*)" with above product url in body$`, s.iSendRequestToWithAboveProductUrlInBody)
	ctx.Step(`^the product url "([^"]*)"$`, s.theProductUrl)
	ctx.Step(`^the response code should be (\d+)$`, s.theResponseCodeShouldBe)
}
