package eclipse

import "fmt"

// ConstraintError represents a custom error for a contstraint things.
type ConstraintError string

func (e ConstraintError) Error() string {
	return string(e)
}

// ConstraintErrorf constructs ConstraintError with formatted message.
func ConstraintErrorf(format string, a ...interface{}) ConstraintError {
	return ConstraintError(fmt.Sprintf(format, a...))
}

// NotFoundError represents a custom error for not found.
type NotFoundError string

func (e NotFoundError) Error() string {
	return string(e)
}

// NotFoundErrorf constructs NotFoundError with formatted message.
func NotFoundErrorf(format string, a ...interface{}) NotFoundError {
	return NotFoundError(fmt.Sprintf(format, a...))
}

// ConflictError represents a custom error for conflict.
type ConflictError string

func (e ConflictError) Error() string {
	return string(e)
}

// ConflictErrorf constructs ConflictError with formatted message.
func ConflictErrorf(format string, a ...interface{}) ConflictError {
	return ConflictError(fmt.Sprintf(format, a...))
}
