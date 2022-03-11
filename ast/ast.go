package ast

import "github.com/vangroan/monkey-lang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

var _ Node = &Program{}

type LetStatement struct {
	Token token.Token
	Name  Identfier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identfier struct {
	Token token.Token
	Value string
}

func (id *Identfier) expressionNode()      {}
func (id *Identfier) TokenLiteral() string { return id.Token.Literal }
