package collections

import (
	"io"
	"strings"

	"github.com/gebv/go-dice/parse"
)

var listGenerators = map[string]Generator{}
var listGeneratorConfigs = map[string]GeneratorConfiger{}

// Reg registration generator by name.
func Reg(g Generator, gc GeneratorConfiger) {
	name := strings.ToLower(strings.TrimSpace(g.Name()))
	listGenerators[name] = g
	listGeneratorConfigs[name] = gc
}

// Get return generator by name.
func Get(name string) (Generator, GeneratorConfiger) {
	name = strings.ToLower(strings.TrimSpace(name))
	return listGenerators[name], listGeneratorConfigs[name]
}

type Generator interface {
	Name() string
	Match(node parse.Node) bool
	Gen(w io.Writer, cfg GeneratorConfig, node parse.Node) error
}

type GeneratorConfig interface{}

type GeneratorConfiger func() GeneratorConfig
