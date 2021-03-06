package ast

import (
	"github.com/geode-lang/geode/pkg/lexer"
	"github.com/geode-lang/geode/pkg/util/log"
)

func (p *Parser) parseVariableDefn(allowDefn bool) VariableDefnNode {
	n := VariableDefnNode{}
	n.Token = p.token
	n.NodeType = nodeVariableDecl
	n.TokenReference.Token = p.token

	if p.atType() {
		n.Type = p.parseType()

		if p.token.Is(lexer.TokIdent) {
			n.Name = NewNamedReference(p.token.Value)
			p.next()
		} else if p.token.Is(lexer.TokAssignment) {

		} else {
			n.SyntaxError()
			log.Fatal("Invalid variable declaration\n")
		}

	} else {
		p.token.SyntaxError()
		log.Fatal("Invalid variable declaration")
	}

	if p.token.Is(lexer.TokAssignment) {
		if allowDefn {
			n.HasValue = true
			p.next()
			n.Body = p.parseExpression()
		} else {
			log.Fatal("Variable Initialization of '%s' is not allowed in it's context\n", n.Name)
		}
	}

	return n
}
