// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package response

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SuccessOk-1]
	_ = x[ErrFail-200000]
	_ = x[ErrInvalidParams-200001]
	_ = x[ErrNoAccess-200002]
	_ = x[ErrNoFound-200003]
	_ = x[ErrDb-200004]
	_ = x[ErrCache-200005]
	_ = x[ErrLimitExceeded-200006]
}

const (
	_ErrCode_name_0 = "ok"
	_ErrCode_name_1 = "服务器错误非法请求参数无访问权限找不到资源数据库出错缓存出错请求国泰，请稍后重试"
)

var (
	_ErrCode_index_1 = [...]uint8{0, 15, 33, 48, 63, 78, 90, 120}
)

func (i ErrCode) String() string {
	switch {
	case i == 1:
		return _ErrCode_name_0
	case 200000 <= i && i <= 200006:
		i -= 200000
		return _ErrCode_name_1[_ErrCode_index_1[i]:_ErrCode_index_1[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
