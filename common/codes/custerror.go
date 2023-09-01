package codes

import (
	"github.com/baowk/dilu-core/core/errs"
)

func ErrSys(cause error) errs.IError {
	return errs.Err(FAILURE, "", cause)
}

func ErrInvalidParameter(reqId string, cause error) errs.IError {
	return errs.Err(InvalidParameter, reqId, cause)
}

func ErrNotFound(id, kind, reqId string, cause error) errs.IError {
	data := map[string]interface{}{"kind": kind, "id": id}
	return errs.ErrWithData(NotFound, reqId, cause, data)
}
