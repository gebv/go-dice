package gen

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Comments print comments with adds for every line "// ".
func Comments(w io.Writer, comments string) {
	scanner := bufio.NewScanner(strings.NewReader(comments))
	for scanner.Scan() {
		fmt.Fprint(w, "// ", scanner.Text())
		fmt.Fprintln(w)
	}
}

// Package print package line with comments.
func Package(f io.Writer, pkgName string, comments ...string) {
	if len(comments) > 0 {
		Comments(f, comments[0])
	}
	fmt.Fprintln(f, "package "+pkgName)
	fmt.Fprintln(f)
}
