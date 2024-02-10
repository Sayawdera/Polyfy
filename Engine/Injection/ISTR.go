package Injection

import (
	"github.com/ddosify/go-faker/faker"
	"regexp"
	"sync"
)

type BodyPiece struct {
	start      int
	end        int
	injectable bool
	value      string
}

type PolygonBodyReader struct {
	Body   string
	Pieces []BodyPiece
	pi     int
	vi     int
}

type EnvironmentInjector struct {
	r   *regexp.Regexp
	jr  *regexp.Regexp
	dr  *regexp.Regexp
	jdr *regexp.Regexp
	mu  sync.Mutex
}

type EnvMatch struct {
	regex string
	found []int
}
type EnvMatchSlice []EnvMatch

var dynamicFakeDataMap map[string]interface{}
var dataFaker faker.Faker
