package servers

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type builder struct {
	str string
}

func createBuilder() Builder {
	out := builder{
		str: "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithURL adds a url to the builder
func (app *builder) WithURL(str string) Builder {
	app.str = str
	return app
}

// Now builds a new Server instance
func (app *builder) Now() (Server, error) {
	if app.str == "" {
		return nil, errors.New("the url is mandatory in order to build a Server instance")
	}

	parsed, err := url.Parse(app.str)
	if err != nil {
		return nil, err
	}

	var host Host
	hostName := parsed.Hostname()
	if strings.HasPrefix(hostName, onionTLD) {
		host = createHostWithOnion(hostName)
	} else {
		host = createHostWithClear(hostName)
	}

	portAsStr := parsed.Port()
	port, err := strconv.Atoi(portAsStr)
	if err != nil {
		return nil, err
	}

	return createServer(parsed.Scheme, host, uint(port)), nil

}
