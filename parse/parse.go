// Package parse for parse golang code.
// Extraction of basic information to generate helpfull code.
package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

// Parse return interested nodes of file.
func Parse(pathFile string) []Node {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(
		fset,
		pathFile,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		log.Fatal("failed parse file", err)
	}

	// ast.Print(fset, f)

	res := []Node{}

	var node Node
	for _, decl := range f.Decls {
		node = Node{}

		switch dt := decl.(type) {
		// import, const, type, var
		case *ast.GenDecl:
			node.NodeDocComments = dt.Doc
			node.TokenType = dt.Tok
			node.NodePosition = fset.Position(dt.TokPos)

			for _, spec := range dt.Specs {

				switch st := spec.(type) {
				case *ast.TypeSpec:
					// type <name> ...
					node.TypeSpec = st
					node.TypeSpecType = st.Type

					node.FieldDocComments = st.Doc
					node.FieldLineComments = st.Comment
					node.FieldName = st.Name.String()

				case *ast.ValueSpec:
					// var, const
					node.ValueSpec = st
					node.ValueSpecType = st.Type

					node.FieldDocComments = st.Doc
					node.FieldLineComments = st.Comment
					node.FieldName = st.Names[0].String()

					if st.Values != nil {
						node.FieldValue = st.Values[0]
					}
				}

				res = append(res, node)
			}
		case *ast.FuncDecl:
			// TODO: not supported
		}
	}

	return res
}
