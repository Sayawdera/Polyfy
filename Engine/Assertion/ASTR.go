package Assertion

import (
	"Polyfy/Engine/Types"
	"Polyfy/Temperature/Evaluator"
	"sync"
)

type AssertionError struct {
	failedAssertion string
	received        map[string]interface{}
	wrappedErr      error
}

type DefaultAssertionService struct {
	assertions map[string]Types.TestAssertionOpt // Rule -> Opts
	abortChan  chan struct{}
	doneChan   chan struct{}
	resChan    chan TestAssertionResult
	assertEnv  *Evaluator.AssertEnv
	abortTick  map[string]int // rule -> tickIndex
	iterCount  int
	mu         sync.Mutex
}

type TestAssertionResult struct {
	Fail        bool         `json:"fail"`
	Aborted     bool         `json:"aborted"`
	FailedRules []FailedRule `json:"failed_rules"`
}

type FailedRule struct {
	Rule        string                 `json:"rule"`
	ReceivedMap map[string]interface{} `json:"received"`
}

var tickerInterval = 100
