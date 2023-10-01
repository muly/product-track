package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/cucumber/godog"
)

type scenarioData struct {
	productUrl string
	statusCode int
	actual     product
	apiHost    string
}

func (s *scenarioData) test(name string) error {
	fmt.Printf("running test %s\n", name)
	return nil
}

func (s *scenarioData) theDeployedApiHost(url string) error {
	s.apiHost = url
	return nil
}

func (s *scenarioData) iSendRequestToWithAboveProductUrlInBody(method, endpoint string) error {
	appURL, err := url.JoinPath(s.apiHost, endpoint)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(strings.ToUpper(method), appURL, strings.NewReader(fmt.Sprintf(`{"url":"%s"}`, s.productUrl)))
	if err != nil {
		return err
	}
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	if response.StatusCode == 200 {
		if err = json.NewDecoder(response.Body).Decode(&s.actual); err != nil {
			return err
		}
	}
	s.statusCode = response.StatusCode
	return nil
}
func (s *scenarioData) theProductUrl(mockProductURL string) error {
	var err error
	s.productUrl, err = url.JoinPath(s.apiHost, mockProductURL)
	if err != nil {
		return err
	}

	return nil
}

func (s *scenarioData) theResponseShouldBe(responseBodyFile string) error {
	if responseBodyFile == "" {
		return nil
	}

	//open the file
	data, err := os.ReadFile("integration_testing/data/" + responseBodyFile)
	if err != nil {
		log.Printf("error during response body %s: %v", responseBodyFile, err)
		return err
	}

	// replace app hostname in the expected
	expectedResponse := strings.Replace(string(data), "{{API_HOST}}", s.apiHost, -1)

	// read the expected body
	var expectedBody product
	if err = json.Unmarshal([]byte(expectedResponse), &expectedBody); err != nil {
		return err
	}

	// compare the expected body and response body
	if expectedBody != s.actual {
		return fmt.Errorf("%+v is not equal to %+v ", expectedBody, s.actual)
	}
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
	ctx.Step(`^test "([^"]*)"$`, s.test)
	ctx.Step(`^the deployed api host "([^"]*)"$`, s.theDeployedApiHost)
	ctx.Step(`^the product url "([^"]*)"$`, s.theProductUrl)
	ctx.Step(`^i send "([^"]*)" request to "([^"]*)" with above product url in body$`, s.iSendRequestToWithAboveProductUrlInBody)
	ctx.Step(`^the response should be "([^"]*)"$`, s.theResponseShouldBe)
	ctx.Step(`^the response code should be (\d+)$`, s.theResponseCodeShouldBe)
}
