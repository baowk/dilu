package codes

import (
	"github.com/baowk/dilu-core/core/errs"
)

func ErrSys(cause error) errs.IError {
	return errs.Err(FAILURE, "", cause)
}

func Err401(cause error) errs.IError {
	return errs.Err(InvalidToken_401, "", cause)
}

func Err403(cause error) errs.IError {
	return errs.Err(AuthorizationError_403, "", cause)
}

// func ErrInvalidParameter(reqId string, cause error) errs.IError {
// 	return errs.Err(InvalidParameter, reqId, cause)
// }

func ErrInvalidParameter(reqId string, msg string) errs.IError {
	data := map[string]interface{}{"msg": msg}
	return errs.ErrWithData(InvalidParameter, reqId, nil, data)
}

func ErrNotFound(id, kind, reqId string, cause error) errs.IError {
	data := map[string]interface{}{"kind": kind, "id": id}
	return errs.ErrWithData(NotFound_404, reqId, cause, data)
}
