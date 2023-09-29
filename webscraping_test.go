package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cucumber/godog"
)

type scenarioData struct {
	productUrl string
	statusCode int 
	actual product
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
	if err= json.NewDecoder(response.Body).Decode(&s.actual) ; err!=nil{
		return err 
	}
	log.Println("actual response:",s.actual)
	s.statusCode = response.StatusCode
	return nil
}
func (s *scenarioData) theProductUrl(url string) error {
	s.productUrl = url
	return nil
}


func (s *scenarioData) theResponseShouldBe(responseBodyFile string) error {

	//open the file
	file, err := os.Open("integration_testing/data/"+responseBodyFile)
	if err != nil {
		log.Printf("error during response body %s: %v",responseBodyFile,err)
		return err 
	}

	//read the file
	var expectedBody product
	if err=json.NewDecoder(file).Decode(&expectedBody);err!=nil{
		return err 
	}
	log.Println("expected response:",expectedBody);
	//bring  the data in the file in json format
	//close the file
	// compare the expected body and response body
	if expectedBody!= s.actual{
		return fmt.Errorf("%+v is not equal to %+v ",expectedBody,s.actual)
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
	ctx.Step(`^i send "([^"]*)" request to "([^"]*)" with above product url in body$`, s.iSendRequestToWithAboveProductUrlInBody)
	ctx.Step(`^the product url "([^"]*)"$`, s.theProductUrl)
	ctx.Step(`^the response should be "([^"]*)"$`, s.theResponseShouldBe)
	ctx.Step(`^the response code should be (\d+)$`, s.theResponseCodeShouldBe)
}
