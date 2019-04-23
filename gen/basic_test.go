package gen

import (
	"bytes"
	"testing"
)

func Test_FormatComments(t *testing.T) {
	tests := []struct {
		name     string
		comments string
		want     string
	}{
		{
			"",
			"",
			"",
		},
		{
			"",
			"foo bar",
			`// foo bar
`,
		},
		{
			"",
			`line1
line 2`,
			"// line1\n// line 2\n",
		},
		{
			"",
			`line1
line 2




`,
			"// line1\n// line 2\n// \n// \n// \n// \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buff := new(bytes.Buffer)
			Comments(buff, tt.comments)
			if buff.String() != tt.want {
				t.Errorf("Comments() = %q, want %q", buff.String(), tt.want)
			}
		})
	}
}
