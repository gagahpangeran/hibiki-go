package hibiki

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Registry Stores list of commands
type Registry struct {
	commands map[string]*Command
}

// DefaultRegistry is the global default command registry
var DefaultRegistry = NewRegistry()

// NewRegistry creates a new CommandRegistry
func NewRegistry() *Registry {
	return &Registry{
		make(map[string]*Command),
	}
}

// AddCommand Add a new command to registry
func (reg *Registry) AddCommand(cmd *Command) {
	_, ok := reg.commands[cmd.identifier]

	if ok {
		logrus.Fatalf(
			"Multiple %#v route declared",
			cmd.identifier,
		)
	}

	cmd.checkValid()
	reg.commands[cmd.identifier] = cmd
}

func (reg *Registry) MatchPrefix(text string) (*Command, *CommandRoute, string) {
	var (
		maxPrefixCommand       *Command
		maxPrefixCommandLength = -1
	)

	text = text + " "

	for prefix, cmd := range reg.commands {
		prefix = prefix + " "
		if strings.HasPrefix(text, prefix) && len(prefix) > maxPrefixCommandLength {
			maxPrefixCommand = cmd
			maxPrefixCommandLength = len(prefix)
		}
	}

	if maxPrefixCommand == nil {
		return nil, nil, ""
	}

	text = text[len(maxPrefixCommand.identifier)+1:]

	route, stringRest := maxPrefixCommand.MatchPrefix(text)

	return maxPrefixCommand, route, stringRest
}

func (reg *Registry) String() string {
	var s strings.Builder
	s.WriteString("Registry { ")

	for _, cmd := range reg.commands {
		s.WriteString(cmd.String())
		s.WriteString(" ")
	}

	s.WriteRune('}')
	return s.String()
}
