// generated by stringer -type ParamType; DO NOT EDIT

package query

import "fmt"

const _ParamType_name = "EmptyIntStringTuple"

var _ParamType_index = [...]uint8{0, 5, 8, 14, 19}

func (i ParamType) String() string {
	if i < 0 || i+1 >= ParamType(len(_ParamType_index)) {
		return fmt.Sprintf("ParamType(%d)", i)
	}
	return _ParamType_name[_ParamType_index[i]:_ParamType_index[i+1]]
}
