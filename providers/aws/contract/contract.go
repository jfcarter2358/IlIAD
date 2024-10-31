package contract

type Comm struct {
	Body       []byte
	StatusCode int64
	Headers    map[string]string
}
