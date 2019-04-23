package filepathx

import "testing"

func Test_FileNameWithSuffix(t *testing.T) {
	tests := []struct {
		pathFile string
		suffix   string
		want     string
	}{
		// basic
		{
			"foo.bar",
			"abc",
			"foo_abc.bar",
		},
		{
			"/path/path/foo.bar",
			"abc",
			"/path/path/foo_abc.bar",
		},
		{
			"/foo.bar",
			"abc",
			"/foo_abc.bar",
		},
		{
			"./foo.bar",
			"abc",
			"foo_abc.bar",
		},

		// without extenstion
		{
			"foo",
			"abc",
			"foo_abc",
		},
		{
			"/path/path/foo",
			"abc",
			"/path/path/foo_abc",
		},
		{
			"/foo",
			"abc",
			"/foo_abc",
		},
		{
			"./foo",
			"abc",
			"foo_abc",
		},

		// empty
		{
			"",
			"abc",
			"",
		},

		// special
		{
			"/path/path/foo.ext1.ext2.ext3",
			"abc",
			"/path/path/foo.ext1.ext2_abc.ext3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.pathFile, func(t *testing.T) {
			if got := FileNameWithSuffix(tt.pathFile, tt.suffix); got != tt.want {
				t.Errorf("fileNameWithSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
