// Code generated by "stringer -type=ManglePartType"; DO NOT EDIT.

package ast

import "strconv"

const _ManglePartType_name = "NamespaceMangleNameMangleGenericMangle"

var _ManglePartType_index = [...]uint8{0, 15, 25, 38}

func (i ManglePartType) String() string {
	if i < 0 || i >= ManglePartType(len(_ManglePartType_index)-1) {
		return "ManglePartType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ManglePartType_name[_ManglePartType_index[i]:_ManglePartType_index[i+1]]
}
