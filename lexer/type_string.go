// generated by stringer -type Type; DO NOT EDIT

package lexer // import "sevki.org/graphql/lexer"

import "fmt"

const _Type_name = "EOFErrorNewlineStringSpaceNumberLeftCurlyRightCurlyLeftParenRightParenLeftBracRightBracQuoteColonCommaSemicolonPeriodCommentVariable"

var _Type_index = [...]uint8{0, 3, 8, 15, 21, 26, 32, 41, 51, 60, 70, 78, 87, 92, 97, 102, 111, 117, 124, 132}

func (i Type) String() string {
	if i < 0 || i+1 >= Type(len(_Type_index)) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
