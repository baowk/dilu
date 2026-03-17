package codes

import (
	"fmt"

	"github.com/baowk/dilu-core/core/errs"
)

// AppError 是可被 CustomError 中间件捕获的类型化错误，用于替代 panic("CustomError#code#msg") 模式
type AppError struct {
	Code int
	Msg  string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("AppError(%d): %s", e.Code, e.Msg)
}

// NewAppError 创建一个 AppError，用于 panic 抛出后被中间件捕获
func NewAppError(code int, msg string) *AppError {
	return &AppError{Code: code, Msg: msg}
}

// PanicApp 抛出 AppError panic，替代原有的 panic("CustomError#code#msg")
func PanicApp(code int, msg string) {
	panic(NewAppError(code, msg))
}

func ErrSys(cause error) errs.IError {
	return errs.Err(FAILURE, "", cause)
}

func Err401(cause error) errs.IError {
	return errs.Err(InvalidToken_401, "", cause)
}

func Err403(cause error) errs.IError {
	return errs.Err(AuthorizationError_403, "", cause)
}

func ErrInvalidParameter(reqId string, msg string) errs.IError {
	data := map[string]interface{}{"msg": msg}
	return errs.ErrWithData(InvalidParameter, reqId, nil, data)
}

func ErrNotFound(id, kind, reqId string, cause error) errs.IError {
	data := map[string]interface{}{"kind": kind, "id": id}
	return errs.ErrWithData(NotFound_404, reqId, cause, data)
}
