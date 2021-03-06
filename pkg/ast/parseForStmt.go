package ast

import (
	"github.com/geode-lang/geode/pkg/lexer"
)

var forStmtIndex = 0

func (p *Parser) parseForStmt() Node {
	p.requires(lexer.TokFor)
	n := ForNode{}
	n.TokenReference.Token = p.token
	n.NodeType = nodeFor
	n.Index = forStmtIndex
	forStmtIndex++
	p.next()

	n.Init = p.parseIdentifierExpr(true)
	p.requires(lexer.TokSemiColon)

	p.next()

	n.Cond = p.parseExpression()
	p.requires(lexer.TokSemiColon)
	p.next()

	n.Step = p.parseExpression()

	p.requires(lexer.TokLeftCurly)
	n.Body = p.parseBlockStmt()

	return n
}
