package errors

import "fmt"

type InternalError struct {
    wError
}

func NewInternalError(msg string) InternalError {
    return InternalError{
        wError{
            retryable: true,
            code:      InternalErrorCode,
            message:   fmt.Sprintf("internal error: %s", msg),
        },
    }
}

type ResourceNotFoundError struct {
    wError
}

func NewResourceNotFoundError(msg string) ResourceNotFoundError {
    return ResourceNotFoundError{
        wError{
            retryable: false,
            code:      ResourceNotFoundErrorCode,
            message:   fmt.Sprintf("resource not found error: %s", msg),
        },
    }
}

type TimeoutError struct {
    wError
}

func NewTimeoutError(msg string) TimeoutError {
    return TimeoutError{
        wError{
            retryable: true,
            code:      TimeoutErrorCode,
            message:   fmt.Sprintf("timeout error: %s", msg),
        },
    }
}

type ValidationError struct {
    wError
}

func NewValidationError(msg string) ValidationError {
    return ValidationError{
        wError{
            retryable: true,
            code:      ValidationErrorCode,
            message:   fmt.Sprintf("validation error: %s", msg),
        },
    }
}

type DuplicatedResourceIdError struct {
    wError
}

func NewDuplicatedResourceIdError(msg string) DuplicatedResourceIdError {
    return DuplicatedResourceIdError{
        wError{
            retryable: true,
            code:      DuplicatedResourceIdErrorCode,
            message:   fmt.Sprintf("duplicated resource id error: %s", msg),
        },
    }
}

type WrongResourceVersionError struct {
    wError
}

func NewWrongResourceVersionError(msg string) WrongResourceVersionError {
    return WrongResourceVersionError{
        wError{
            retryable: true,
            code:      WrongResourceVersionErrorCode,
            message:   fmt.Sprintf("wrong resource version error: %s", msg),
        },
    }
}
