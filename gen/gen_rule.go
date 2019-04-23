package gen

import (
	"io"
	"log"
	"strings"

	// toml "github.com/pelletier/go-toml"
	toml "github.com/BurntSushi/toml"
	"github.com/gebv/go-dice/gen/collections"
	"github.com/gebv/go-dice/parse"
)

type GeneratorRule struct {
	AnnotaionName string
	Cfg           collections.GeneratorConfig
	CfgRaw        string
	G             collections.Generator
}

func NewGeneratorRule(a Annotation) *GeneratorRule {
	r := &GeneratorRule{
		AnnotaionName: a.Name,
		CfgRaw:        a.Body,
	}
	r.setup()
	return r
}

func (r *GeneratorRule) Name() string {
	return r.G.Name()
}

// func (r *GeneratorRule) Match(t token.Token, spec ast.Spec) bool {
func (r *GeneratorRule) Match(node parse.Node) bool {
	return r.G.Match(node)
}

func (r *GeneratorRule) Gen(w io.Writer, cfg collections.GeneratorConfig, node parse.Node) error {
	return r.G.Gen(w, cfg, node)
}

var _ collections.Generator = (*GeneratorRule)(nil)

func (r *GeneratorRule) IsValid() bool {
	return r.Cfg != nil && r.G != nil
}

func (r *GeneratorRule) setup() {
	g, gc := collections.Get(strings.ToLower(strings.TrimSpace(r.AnnotaionName)))
	if g == nil && gc == nil {
		log.Printf("Not found generator %q - skip", r.AnnotaionName)
		return
	}
	r.Cfg = gc()

	if err := toml.Unmarshal([]byte(r.CfgRaw), r.Cfg); err != nil {
		log.Fatalf("Failed decode annotation %q - %v", r.AnnotaionName, err)
		return
	}
	r.G = g
}
