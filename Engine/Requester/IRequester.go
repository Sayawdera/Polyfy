package Requester

import (
	"Polyfy/Engine/Types"
	"context"
	"net/http"
	"net/url"
)

type Requester interface {
	Type() string
	Done()
}

type HttpRequesterI interface {
	Init(ctx context.Context, ss Types.ScenarioStep, url *url.URL, debug bool, ei *Injection.EnvironmentInjector) error
	Send(client *http.Client, envs map[string]interface{}) *Types.ScenarioStepResult
}
