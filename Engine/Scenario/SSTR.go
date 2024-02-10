package Scenario

import (
	"Polyfy/Engine/Types"
	"Polyfy/Engine/Utils"
	"context"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
)

type ScenarioService struct {
	clients     map[*url.URL][]scenarioItemRequester
	cPool       *Utils.Pool[*http.Client]
	scenario    Types.Scenario
	ctx         context.Context
	clientMutex sync.Mutex
	debug       bool
	engineMode  string
	ei          *injection.EnvironmentInjector
	iterIndex   int64
}

type cookieJarRepeated struct {
	defaultCookieJar *cookiejar.Jar
	firstIterPassed  bool
}
type ScenarioOpts struct {
	Debug                  bool
	IterationCount         int
	MaxConcurrentIterCount int
	EngineMode             string
	InitialCookies         []*http.Cookie
}

type scenarioItemRequester struct {
	scenarioItemID uint16
	sleeper        Sleeper
	requester      Requester.Requester
}

type DurationSleep struct {
	duration int
}
type RangeSleep struct {
	min int
	max int
}

var defaultFactory = func() *http.Client {
	return &http.Client{}
}

var defaultClose = func(c *http.Client) {
	c.CloseIdleConnections()
}
