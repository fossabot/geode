package lexer

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/geode-lang/geode/pkg/typesystem"
	"github.com/geode-lang/geode/pkg/util/color"
)

// TokenIsOperator will return if a given token is an operator or not
func TokenIsOperator(t TokenType) bool {
	return t > TokOperatorStart && t < TokOperatorEnd
}

// Token is a token in the program
type Token struct {
	source *Sourcefile
	Type   TokenType `json:"type,omitempty"`
	Value  string    `json:"value,omitempty"`
	Pos    int       `json:"start_pos"`
	EndPos int       `json:"end_pos"`
	Line   int       `json:"line"`
	Column int       `json:"column"`

	SpaceBefore bool `json:"space_before"`
	SpaceAfter  bool `json:"space_after"`
}

// Is - returns if the given token is in the set of types given
func (t Token) Is(types ...TokenType) bool {
	for _, a := range types {
		if t.Type == a {
			return true
		}
	}
	return false
}

func (t Token) String() string {
	return fmt.Sprintf("%s(%q)", t.Type.String(), t.Value)
}

// SyntaxError prints a formatted syntax error
func (t *Token) SyntaxError() {
	if t.Type == TokError {
		return
	}
	buf := &bytes.Buffer{}
	src := t.source.String()

	// Highlight the source string at the error
	src = src[:t.Pos] + color.Red(src[t.Pos:t.EndPos]) + src[t.EndPos:]
	// Replace tabs with a fixed number of spaces
	src = strings.Replace(src, "\t", "    ", -1)
	lines := strings.Split(src, "\n")

	location := fmt.Sprintf("%s:%d:%d-%d", t.source.Path, t.Line, t.Column, t.Column+len(t.Value))
	// Start printing
	fmt.Fprintf(buf, "\nSyntax error: (%s)\n", location)
	fmt.Fprintf(buf, color.Blue("   |\n"))

	lineNumber := color.Red(fmt.Sprintf("%2d", t.Line))
	fmt.Fprintf(buf, "%s %s %s\n", lineNumber, color.Blue("|"), strings.TrimSpace(lines[t.Line-1]))
	fmt.Fprintf(buf, color.Blue("   |\n"))

	fmt.Println(buf)
}

// InferType takes some token and guesses the type
func (t Token) InferType() (*typesystem.VarType, interface{}) {
	if t.Type == TokNumber {
		intval, intErr := strconv.ParseInt(t.Value, 10, 64)
		if intErr == nil {
			return typesystem.GeodeI64, intval
		}

		floatval, floatErr := strconv.ParseFloat(t.Value, 64)
		if floatErr == nil {
			return typesystem.GeodeF64, floatval
		}
	}

	if t.Type == TokChar {
		c := strings.Trim(t.Value, "'")[0]
		return typesystem.GeodeI8, c
	}

	return nil, nil
}
