package errors

import (
	"net/http"
)

// ErrCode ...
type ErrCode int

const (
	// StatusUnknown ...
	StatusOK           ErrCode = 0
	StatusAlreadyExist ErrCode = 1

	StatusBadRequest       ErrCode = 10000
	StatusUnauthorized     ErrCode = 10001
	StatusForbidden        ErrCode = 10003
	StatusNotFound         ErrCode = 10004
	StatusMethodNotAllowed ErrCode = 10005
	StatusConflict         ErrCode = 10009

	StatusParamsMissing        ErrCode = 10010
	StatusInvalidParams        ErrCode = 10011
	StatusRequireLength        ErrCode = 10012
	StatusEntityTooLarge       ErrCode = 10013
	StatusOutOfRange           ErrCode = 10014
	StatusUnsupportedMediaType ErrCode = 10015
	StatusTooManyRequests      ErrCode = 10029

	ResouceInsufficient ErrCode = 12000
	DataNoAllowOperate  ErrCode = 12001

	StatusAuthInvalidSession ErrCode = 19000
	StatusAuthInvalidCode    ErrCode = 19001
	StatusAuthNoUserInfo     ErrCode = 19002
	StatusAuthInvalidUser    ErrCode = 19003
	StatusInvalidToken       ErrCode = 19004
	StatusNotImplemented     ErrCode = 19005

	StatusInternalDBError     ErrCode = 26000
	StatusInternalServerError ErrCode = 27000
	DataError                 ErrCode = 28000
	StatusUnknown             ErrCode = 99999
)

// Error used to transfer server error
type Error struct {
	error
	errCode ErrCode
}

// NewError define new err type
func NewError(code ErrCode, err error) error {
	return &Error{err, code}
}

// ErrorCode return code
func (e *Error) ErrorCode() ErrCode {
	return e.errCode
}

/*
	ToHTTPCode transfer inner code to http code

Doc: https://docs.google.com/document/d/1hHn65eFDrlwjmYVcFPzZLzLX71Z8pTs3M4zP4utdQIQ/edit
Table:

	00000: http.StatusOk == 200
	10000: http.StatusBadRequest == 400
	10001: http.StatusUnauthorized == 401
	10003: http.StatusForbidden == 403
	10004: http.StatusForbidden == 404
	10005: http.StatusMethodNotAllowed == 405
	10009: http.StatusConflict == 409
	10029: http.StatusTooManyRequests == 429
	[10060, 20000): http.StatusBadRequest == 400
	[20000, 99999): http.StatusInternalServerError == 500
*/
func (code ErrCode) ToHTTPCode() int {
	v := int(code)
	switch {
	case v == 0:
		return http.StatusOK
	case v == 1:
		return http.StatusCreated
	case v == 10000:
		return http.StatusBadRequest
	case v == 10001:
		return http.StatusUnauthorized
	case v == 10003:
		return http.StatusForbidden
	case v == 10004:
		return http.StatusNotFound
	case v == 10005:
		return http.StatusMethodNotAllowed
	case v == 10009:
		return http.StatusConflict
	case v == 10010 || v == 10011:
		return http.StatusBadRequest
	case v == 10012:
		return http.StatusLengthRequired
	case v == 10013:
		return http.StatusRequestEntityTooLarge
	case v == 10014:
		return http.StatusRequestedRangeNotSatisfiable
	case v == 10015:
		return http.StatusUnsupportedMediaType
	case v == 10029:
		return http.StatusTooManyRequests
	case v == 19004:
		return http.StatusBadRequest
	case v == 19005:
		return http.StatusNotImplemented
	case v >= 10060 && v < 20000:
		return http.StatusBadRequest
	case v > 20000:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func StatusUnauthorizedError(err error) *Error {
	return &Error{err, StatusUnauthorized}
}

func ParamsMissingError(err error) *Error {
	return &Error{err, StatusParamsMissing}
}

func InvalidParamsError(err error) *Error {
	return &Error{err, StatusInvalidParams}
}

func ResouceInsufficientError(err error) *Error {
	return &Error{err, ResouceInsufficient}
}

func DataNoAllowOperateError(err error) *Error {
	return &Error{err, DataNoAllowOperate}
}

func InternalDBError(err error) *Error {
	return &Error{err, StatusInternalDBError}
}

func InternalServerError(err error) *Error {
	return &Error{err, StatusInternalServerError}
}

func DataServerError(err error) *Error {
	return &Error{err, DataError}
}

func NotFoundError(err error) *Error {
	return &Error{err, StatusNotFound}
}

func MethodNotAllowedError(err error) *Error {
	return &Error{err, StatusMethodNotAllowed}
}

func ConflictError(err error) *Error {
	return &Error{err, StatusConflict}
}

func ForbiddenError(err error) *Error {
	return &Error{err, StatusForbidden}
}

func AlreadyExistError(err error) *Error {
	return &Error{err, StatusAlreadyExist}
}

func NotImplementedError(err error) *Error {
	return &Error{err, StatusNotImplemented}
}

func InvalidTokenError(err error) *Error {
	return &Error{err, StatusInvalidToken}
}

func OutOfRangeError(err error) *Error {
	return &Error{err, StatusOutOfRange}
}

func EntityTooLargeError(err error) *Error {
	return &Error{err, StatusEntityTooLarge}
}

func LengthRequiredError(err error) *Error {
	return &Error{err, StatusRequireLength}
}

func UnSupportedMediaTypeError(err error) *Error {
	return &Error{err, StatusUnsupportedMediaType}
}

func TooManyRequestsError(err error) *Error {
	return &Error{err, StatusTooManyRequests}
}
