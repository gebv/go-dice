package gen

import (
	"reflect"
	"testing"
)

func TestParseAnnotations(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []Annotation
	}{
		{},
		{
			"",
			"@dice()",
			[]Annotation{
				{Name: "dice"},
			},
		},
		{
			"",
			"@dice(foobar)",
			[]Annotation{
				{Name: "dice", Body: "foobar"},
			},
		},
		{
			"",
			`Some comments (foo bar)
Some comments
	@foo(
	content
	of
	@annotations
)
  @bar(content
	of
	annotations)

Some @comment
`,
			[]Annotation{
				{Name: "foo", Body: "\n\tcontent\n\tof\n\tannotations\n"},
				{Name: "bar", Body: "content\n\tof\n\tannotations"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAnnotations(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAnnotations() = %v, want %v", got, tt.want)
			}
		})
	}
}
