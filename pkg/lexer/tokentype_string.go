// Code generated by "stringer -type=TokenType TokenType.go"; DO NOT EDIT.

package lexer

import "strconv"

const _TokenType_name = "TokErrorTokNoEmitTokWhitespaceTokCharTokStringTokNumberTokDotTokElipsisTokOperTokPtrTokNamespaceAccessTokOperatorStartTokStarTokPlusTokMinusTokDivTokExpTokLTTokLTETokGTTokGTETokOperatorEndTokSemiColonTokDefereferenceTokReferenceTokAssignmentTokEqualityTokRightParenTokLeftParenTokRightCurlyTokLeftCurlyTokRightBraceTokLeftBraceTokRightArrowTokLeftArrowTokCompoundAssignmentTokForTokWhileTokIfTokElseTokReturnTokFuncDefnTokClassDefnTokNamespaceTokDependencyTokTypeTokCommaTokIdentTokComment"

var _TokenType_index = [...]uint16{0, 8, 17, 30, 37, 46, 55, 61, 71, 78, 84, 102, 118, 125, 132, 140, 146, 152, 157, 163, 168, 174, 188, 200, 216, 228, 241, 252, 265, 277, 290, 302, 315, 327, 340, 352, 373, 379, 387, 392, 399, 408, 419, 431, 443, 456, 463, 471, 479, 489}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
