package Report

import "Polyfy/Engine/Types"

type ReportService interface {
	DoneChan() <-chan bool
	Init(debug bool, samplingRate int) error
	Start(input chan *Types.ScenarioResult, assertionResultChan <-chan Assertion.TestAssertionResult)
}

var AvailableOutputServices = make(map[string]ReportService)
