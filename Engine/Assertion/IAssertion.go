package Assertion

import (
	"Polyfy/Engine/Types"
)

type IAborter interface {
	AbortChan() <-chan struct{}
}

type IResultListener interface {
	Start(input <-chan *Types.ScenarioResult)
	DoneChan() <-chan struct{} // indicates processing of results are done
}

type IAsserter interface {
	ResultChan() <-chan TestAssertionResult
}
