package parse

import (
	"go/ast"
	"go/token"
)

// Node is backet with the prepared information.
type Node struct {
	TokenType       token.Token
	NodeDocComments *ast.CommentGroup
	NodePosition    token.Position

	FieldName  string
	FieldValue ast.Expr
	// IdentType  ast.Expr

	FieldDocComments  *ast.CommentGroup
	FieldLineComments *ast.CommentGroup

	// var, const
	ValueSpec *ast.ValueSpec
	// value type; or nil
	ValueSpecType ast.Expr

	// type ...
	TypeSpec *ast.TypeSpec
	// *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
	TypeSpecType ast.Expr
}

// IsVar return true if node is var.
func (n Node) IsVar() bool {
	return n.TokenType == token.VAR && n.ValueSpec != nil
}

// IsConst return true if node is const.
func (n Node) IsConst() bool {
	return n.TokenType == token.CONST && n.ValueSpec != nil
}

// IsType return true if node is type.
func (n Node) IsType() bool {
	return n.TokenType == token.TYPE && n.TypeSpec != nil
}

func (n Node) DocComments() string {
	doc := n.FieldDocComments
	if doc == nil {
		doc = n.NodeDocComments
	}
	if doc == nil {
		return ""
	}
	return doc.Text()
}
