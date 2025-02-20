package helper

import (
	"errors"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errorCode = map[error]codes.Code{
	model.ErrServerError:  codes.Internal,
	model.ErrNotFound:     codes.NotFound,
	model.ErrConflict:     codes.FailedPrecondition,
	model.ErrBadLoginPass: codes.Unauthenticated,
	model.ErrNoClaims:     codes.Internal,
	model.ErrBadParams:    codes.InvalidArgument,
	model.ErrAccessDenied: codes.PermissionDenied,
}

// ErrorToGRPC converts the given model error into a gRPC error.
func ErrorToGRPC(err error) error {
	cause := errors.Unwrap(err)
	if cause == nil {
		cause = err
	}

	if code := status.Code(cause); code != codes.Unknown {
		return cause
	}

	code, ok := errorCode[cause]
	if !ok {
		code = codes.Unknown
	}
	return status.Errorf(code, err.Error())
}
