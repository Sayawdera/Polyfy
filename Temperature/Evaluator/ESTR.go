package Evaluator

import "net/http"

type AssertEnv struct {
	StatusCode    int64
	ResponseSize  int64
	ResponseTime  int64
	Body          string
	Headers       http.Header
	Variables     map[string]interface{}
	Cookies       map[string]*http.Cookie
	TotalTime     []int64
	FailCount     int
	FailCountPerc float64
}

var AssertionFuncMap = map[string]struct{}{
	NOT:          {},
	LESSTHAN:     {},
	GREATERTHAN:  {},
	EQUALS:       {},
	EQUALSONFILE: {},
	IN:           {},
	JSONPATH:     {},
	XMLPATH:      {},
	HTMLPATH:     {},
	REGEXP:       {},
	EXISTS:       {},
	CONTAINS:     {},
	RANGE:        {},
	MIN:          {},
	MAX:          {},
	AVG:          {},
	P99:          {},
	P98:          {},
	P95:          {},
	P90:          {},
	P80:          {},
	TIME:         {},
}

const (
	NOT          = "not"
	LESSTHAN     = "less_than"
	GREATERTHAN  = "greater_than"
	EQUALS       = "equals"
	IN           = "in"
	JSONPATH     = "json_path"
	XMLPATH      = "xpath"
	HTMLPATH     = "html_path"
	REGEXP       = "regexp"
	EXISTS       = "exists"
	CONTAINS     = "contains"
	RANGE        = "range"
	EQUALSONFILE = "equals_on_file"
	TIME         = "time"

	MIN = "min"
	MAX = "max"
	AVG = "avg"
	P99 = "p99"
	P98 = "p98"
	P95 = "p95"
	P90 = "p90"
	P80 = "p80"
)
