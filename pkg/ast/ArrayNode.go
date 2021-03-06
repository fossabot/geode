package ast

import (
	"github.com/geode-lang/geode/pkg/util/log"
	"github.com/geode-lang/llvm/ir/constant"
	"github.com/geode-lang/llvm/ir/types"
	"github.com/geode-lang/llvm/ir/value"
)

// ArrayNode -
type ArrayNode struct {
	NodeType
	TokenReference

	Elements []Node
	Length   int
}

// NameString implements Node.NameString
func (n ArrayNode) NameString() string { return "ArrayNode" }

// InferType implements Node.InferType
func (n ArrayNode) InferType(scope *Scope) string { return "void" }

// GenAccess -
func (n ArrayNode) GenAccess(s *Scope, c *Compiler) value.Value {
	return n.Codegen(s, c)
}

// Codegen implements Node.Codegen for ArrayNode
func (n ArrayNode) Codegen(scope *Scope, c *Compiler) value.Value {
	block := c.CurrentBlock()

	var elementType types.Type
	values := make([]value.Value, 0)
	for _, el := range n.Elements {
		val := el.Codegen(scope, c)
		if elementType == nil {
			elementType = val.Type()
		}

		if !types.Equal(val.Type(), elementType) {
			el.SyntaxError()
			log.Fatal("Invalid type in array. Array should be of a single type '%s', got '%s'\n", elementType, val.Type())
		}
		values = append(values, val)
	}
	typ := c.typeCache

	if len(values) > 0 {
		typ = values[0].Type()
	} else {
		typ = typ.(*types.PointerType).Elem
	}

	typ = types.NewArray(typ, int64(n.Length))

	alloca := block.NewAlloca(typ)

	zero := constant.NewInt(int64(0), types.I64)
	one := constant.NewInt(int64(1), types.I64)
	arrayStart := block.NewGetElementPtr(alloca, zero, zero)
	offset := arrayStart

	for i, val := range values {
		block.AppendInst(NewLLVMComment("[%d] <- %s", i, val.Ident()))
		if i > 0 {
			offset = block.NewGetElementPtr(offset, one)
		}

		block.NewStore(val, offset)
	}

	return arrayStart
}
