package error

//go:generate stringer -type ErrCode -output code_string.go -linecomment

type ErrCode int

const (
	ErrCodeOk ErrCode = 0  // OK
	ErrCodeInvalidParams ErrCode = 1 // 无效的参数
	ErrCodeTimeout   ErrCode = 2 // 请求超时
)

