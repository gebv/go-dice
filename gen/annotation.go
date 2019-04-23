package gen

import (
	"strings"
)

// Annotation conatiner with name and body annotation.
type Annotation struct {
	Name string
	Body string
}

// ParseAnnotations return annotation from text.
//
// Annotation format @<annotation_name(<annotation_body>).
func ParseAnnotations(in string) []Annotation {
	if in == "" {
		return nil
	}
	in = strings.TrimSpace(in)
	res := []Annotation{}

	var state = annotationSearch
	var curr Annotation
	for i := range in {

		switch in[i] {
		case '@':
			if state != annotationSearch {
				break
			}
			state = annotationName
		case '(':
			if state != annotationName {
				break
			}
			state = annotationBodyStart
		case ')':
			if state != annotationBodyStart {
				break
			}
			state = annotationBodyEnd
		default:
			switch state {
			case annotationName:
				curr.Name += string(in[i])
			case annotationBodyStart:
				curr.Body += string(in[i])
			}
		}

		switch state {
		case annotationBodyEnd:
			res = append(res, curr)
			curr = Annotation{}
			state = annotationSearch
		}
	}

	return res
}

type parserState int

const (
	annotationSearch parserState = iota
	annotationName
	annotationBodyStart
	annotationBodyEnd
)
