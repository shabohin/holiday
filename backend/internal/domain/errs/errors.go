package errs

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"text/template"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ErrorCode uint

const (
	ErrorCodeOK ErrorCode = iota
	ErrorCodeCanceled
	ErrorCodeUnknown
	ErrorCodeInvalidArgument
	ErrorCodeDeadlineExceeded
	ErrorCodeNotFound
	ErrorCodeAlreadyExists
	ErrorCodePermissionDenied
	ErrorCodeResourceExhausted
	ErrorCodeFailedPrecondition
	ErrorCodeAborted
	ErrorCodeOutOfRange
	ErrorCodeUnimplemented
	ErrorCodeInternal
	ErrorCodeUnavailable
	ErrorCodeDataLoss
	ErrorCodeUnauthenticated
)

type Error struct {
	Code    ErrorCode         `json:"code"`
	Message string            `json:"message"`
	Params  map[string]string `json:"params"`
}

func (e Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e *Error) Is(tgt error) bool {
	target, ok := tgt.(*Error)
	if !ok {
		return false
	}
	return reflect.DeepEqual(e, target)
}

func (e *Error) SetCode(code ErrorCode) {
	e.Code = code
}

func (e *Error) SetMessage(message string) {
	e.Message = message
}

func (e *Error) SetParams(params map[string]string) {
	e.Params = params
}

func (e *Error) AddParam(key string, value string) {
	e.Params[key] = value
}

func NewError(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Params:  map[string]string{},
	}
}

func NewUnexpectedBehaviorError(details string) *Error {
	return &Error{
		Code:    ErrorCodeInternal,
		Message: "Unexpected behavior.",
		Params: map[string]string{
			"details": details,
		},
	}
}

func NewInvalidFormError() *Error {
	return NewError(ErrorCodeInvalidArgument, "The form sent is not valid, please correct the errors below.")
}

func NewInvalidParameter(message string) *Error {
	e := NewError(ErrorCodeInvalidArgument, message)
	return e
}

func FromValidationError(err error) *Error {
	var validationErrors validation.Errors
	var validationErrorObject validation.ErrorObject
	if errors.As(err, &validationErrors) {
		e := NewError(ErrorCodeInvalidArgument, "The form sent is not valid, please correct the errors below.")
		for key, value := range validationErrors {
			switch t := value.(type) {
			case validation.ErrorObject:
				e.AddParam(key, renderErrorMessage(t))
			case *Error:
				e.AddParam(key, t.Message)
			default:
				e.AddParam(key, value.Error())
			}
		}
		return e
	}
	if errors.As(err, &validationErrorObject) {
		return NewInvalidParameter(renderErrorMessage(validationErrorObject))
	}
	return nil
}

func renderErrorMessage(object validation.ErrorObject) string {
	parse, err := template.New("message").Parse(object.Message())
	if err != nil {
		return ""
	}
	var tpl bytes.Buffer
	_ = parse.Execute(&tpl, object.Params())
	return tpl.String()
}

func NewEventNotFound() *Error {
	return &Error{
		Code:    ErrorCodeNotFound,
		Message: "Event not founded",
		Params:  map[string]string{},
	}
}
