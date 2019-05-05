package fsmchannelwait

import (
	"fmt"
	"go/ast"
	"io"

	"github.com/pkg/errors"

	"github.com/gebv/go-dice/gen/collections"
	"github.com/gebv/go-dice/parse"
)

type Config struct{}

func init() {
	collections.Reg(&Generator{}, func() collections.GeneratorConfig {
		// config with default values
		return &Config{}
	})
}

const NAME = "FSMChannelWait"

type Generator struct {
}

func (*Generator) Name() string {
	return NAME
}

// Match return true if matched of node element.
// Interested only type struct.
func (*Generator) Match(node parse.Node) bool {
	if !node.IsType() {
		return false
	}

	_, ok := node.TypeSpec.Type.(*ast.StructType)
	return ok
}

func (t *Generator) Gen(w io.Writer, icfg collections.GeneratorConfig, node parse.Node) error {
	_, ok := icfg.(*Config)
	if !ok {
		return fmt.Errorf("invalid config type - got %T", icfg)
	}

	err := tpl.Execute(w, map[string]interface{}{
		"StructName": node.FieldName,
	})
	if err != nil {
		return errors.Wrap(err, "failed execute template for fsmchannelwait")
	}

	return nil
}

var _ collections.Generator = (*Generator)(nil)
var _ collections.GeneratorConfig = (*Config)(nil)
