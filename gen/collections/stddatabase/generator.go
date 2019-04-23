package stddatabase

import (
	"fmt"
	"go/ast"
	"io"

	"github.com/pkg/errors"

	"github.com/gebv/go-dice/gen/collections"
	"github.com/gebv/go-dice/parse"
)

func init() {
	collections.Reg(&Generator{}, func() collections.GeneratorConfig {
		// config with default values
		return &Config{
			ScanPrint:        true,
			ValuePrint:       true,
			ScanEmptyAsError: false,
		}
	})
}

const NAME = "sqldriver"

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
	cfg, ok := icfg.(*Config)
	if !ok {
		return fmt.Errorf("invalid config type - got %T", icfg)
	}

	if cfg.ScanPrint {
		err := tplSqlDriverScan.Execute(w, map[string]interface{}{
			"StructName":   node.FieldName,
			"EmptyIsError": cfg.ScanEmptyAsError,
		})
		if err != nil {
			return errors.Wrap(err, "failed execute template for sql/database/driver#scan")
		}
	}

	if cfg.ValuePrint {
		err := tplSqlDriverValue.Execute(w, map[string]interface{}{
			"StructName": node.FieldName,
		})
		if err != nil {
			return errors.Wrap(err, "failed execute template for sql/database/driver#value")
		}
	}

	return nil
}

var _ collections.Generator = (*Generator)(nil)
var _ collections.GeneratorConfig = (*Config)(nil)
