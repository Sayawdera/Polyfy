package Data

type RemoteCsvError struct {
	Message    string
	wrappedErr error
}
