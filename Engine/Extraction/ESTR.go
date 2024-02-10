package Extraction

import "regexp"

type ExtractionError struct {
	Message    string
	wrappedErr error
}

type htmlExtractor struct {
}

type jsonExtractor struct {
}

type regexExtractor struct {
	r *regexp.Regexp
}

type xmlExtractor struct {
}
