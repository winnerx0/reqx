package utils

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

type Config struct {
	Requests []Request `yaml:"requests"`
}

type Request struct {
	Name    string            `yaml:"name"`
	Url     string            `yaml:"url"`
	Body    map[string]any    `yaml:"body"`
	Headers map[string]string `yaml:"headers"`
	Method  Method            `yaml:"method"`
}
