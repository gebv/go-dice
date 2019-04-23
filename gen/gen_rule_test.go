package gen

import (
	"io"
	"reflect"
	"testing"

	"github.com/gebv/go-dice/gen/collections"
	_ "github.com/gebv/go-dice/gen/collections/stddatabase"
	"github.com/gebv/go-dice/parse"
)

func init() {
	collections.Reg(&genTest{}, func() collections.GeneratorConfig {
		return &genTestConfig{
			// default falues
			Foo: true,
			Bar: "123",
		}
	})
}

func TestNewGeneratorRule(t *testing.T) {
	tests := []struct {
		name  string
		a     Annotation
		want  *GeneratorRule
		valid bool
	}{
		{
			"empty",
			Annotation{},
			&GeneratorRule{},
			false,
		},
		{
			"notRegGenerator",
			Annotation{Name: "notFound", Body: `b = "b"
			int1 = 123
			float2 = 123.123
			arrInt1 = [1,2,3]
			arrStr1 = ["a", "b", "c"]`},
			&GeneratorRule{
				AnnotaionName: "notFound",
				Cfg:           nil,
			},
			false,
		},
		{
			"regGenCheckDefConfig",
			Annotation{Name: "genTest", Body: `b = "b"
			int1 = 123
			float2 = 123.123
			arrInt1 = [1,2,3]
			arrStr1 = ["a", "b", "c"]`},
			&GeneratorRule{
				AnnotaionName: "genTest",
				Cfg: &genTestConfig{
					Foo: true,
					Bar: "b",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGeneratorRule(tt.a)
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("NewGeneratorRule().Name = %v, want %v", got.AnnotaionName, tt.want.AnnotaionName)
			}
			if !reflect.DeepEqual(got.Cfg, tt.want.Cfg) {
				t.Errorf("NewGeneratorRule().Cfg = %v, want %v", got.Cfg, tt.want.Cfg)
			}
			if got.IsValid() != tt.valid {
				t.Errorf("NewGeneratorRule().IsValid = %v, want %v", got.IsValid(), tt.valid)
			}
		})
	}
}

type genTest struct{}

func (*genTest) Name() string {
	return "genTest"
}

func (*genTest) Match(parse.Node) bool {
	panic("not implemented")
}

func (t *genTest) Gen(w io.Writer, cfg collections.GeneratorConfig, node parse.Node) error {
	panic("not implemented")
}

type genTestConfig struct {
	Foo bool   `toml:"a"`
	Bar string `toml:"b"`
}
