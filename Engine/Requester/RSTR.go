package Requester

import (
	"Polyfy/Engine/Types"
	"context"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

type HttpRequester struct {
	ctx                  context.Context
	proxyAddr            *url.URL
	packet               Types.ScenarioStep
	client               *http.Client
	request              *http.Request
	ei                   *Injection.EnvironmentInjector
	containsDynamicField map[string]bool
	containsEnvVar       map[string]bool
	debug                bool
	dynamicRgx           *regexp.Regexp
	envRgx               *regexp.Regexp
}

type duration struct {
	dnsDur                     time.Duration
	connDur                    time.Duration
	tlsDur                     time.Duration
	reqDur                     time.Duration
	resDur                     time.Duration
	serverProcessDur           time.Duration
	resStart                   time.Time
	resStartCh                 chan time.Time
	resStartChClosed           bool
	serverProcessDurCh         chan time.Duration
	serverProcessDurChClosed   bool
	serverProcessStartCh       chan time.Time
	serverProcessStartChClosed bool
	resDurCh                   chan time.Duration
	resDurChClosed             bool
	mu                         sync.Mutex
	chMu                       sync.Mutex
	getChLock                  sync.Mutex
}

var durationCloseFunc = func(d *duration) func() {
	return func() {
		d.close()
	}
}
