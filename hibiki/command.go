package hibiki

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// Command stores a command and it's routes
type Command struct {
	identifier string
	routes     map[string]*CommandRoute
}

// CommandRoute represents a single command route
type CommandRoute struct {
	identifier string
	definition string
	handler    CommandRouteHandler
}

// CommandRouteHandler function that response to route
type CommandRouteHandler func(req *Request, res *Response) error

// NewCommand Creates new Command object
func NewCommand(identifier string, routes ...*CommandRoute) *Command {
	cmd := &Command{
		strings.TrimSpace(identifier),
		make(map[string]*CommandRoute),
	}
	for _, route := range routes {
		cmd.AddRoute(route)
	}
	return cmd
}

func (c *Command) AddRoute(route *CommandRoute) {
	_, ok := c.routes[route.identifier]
	if ok {
		logrus.Fatalf(
			"Handler for command '%v' have multiple %#v handler",
			c.identifier,
			route.identifier,
		)
	}
	c.routes[route.identifier] = route
}

func (c *Command) MatchPrefix(text string) (*CommandRoute, string) {
	var (
		maxPrefixRoute       *CommandRoute
		maxPrefixRouteLength = -1
	)

	text = strings.TrimSpace(text) + " "

	for prefix, route := range c.routes {
		prefix = prefix + " "
		if strings.HasPrefix(text, prefix) && len(prefix) > maxPrefixRouteLength {
			maxPrefixRoute = route
			maxPrefixRouteLength = len(prefix)
		}
	}

	if maxPrefixRoute == nil {
		maxPrefixRoute = c.routes[""]
	}

	text = text[len(maxPrefixRoute.identifier):]
	text = strings.TrimSpace(text)

	return maxPrefixRoute, text
}

func (c *Command) checkValid() {
	for _, route := range c.routes {
		if route.identifier == "" {
			return
		}
	}

	logrus.Fatalf(
		"Command '%v' missing default route (\"\")",
		c.identifier,
	)
}

func (c *Command) String() string {
	var s strings.Builder

	s.WriteString(c.identifier)
	s.WriteString(" [ ")
	for _, route := range c.routes {
		s.WriteString(fmt.Sprintf("%#v ", route.identifier))
	}
	s.WriteString("]")

	return s.String()
}

// NewDefaultRoute creates a new command route
func NewDefaultRoute(definition string, handler CommandRouteHandler) *CommandRoute {
	return NewCommandRoute("", definition, handler)
}

// NewCommandRoute creates a new command route
func NewCommandRoute(identifier string, definition string, handler CommandRouteHandler) *CommandRoute {
	return &CommandRoute{
		strings.TrimSpace(identifier),
		strings.TrimSpace(definition),
		handler,
	}
}

func (r *CommandRoute) handle(req *Request, res *Response) error {
	return r.handler(req, res)
}
