package filter

import (
	"errors"
	"strconv"
	"strings"
)

// ErrorKind classifies failures returned by this package and its subpackages.
//
// Kinds are deliberately coarse: four buckets cover every failure mode a
// pure value-transformation library can produce. Callers branch on Kind to
// map filter errors onto their own exit codes / HTTP status / log levels
// without maintaining a sentinel translation table.
type ErrorKind uint8

const (
	// KindInvalidInput is returned when the caller passed a value of the
	// wrong type, the wrong shape, or out of range.
	KindInvalidInput ErrorKind = iota + 1
	// KindNotFound is returned when a path, key, or index does not exist
	// in the input.
	KindNotFound
	// KindArithmetic is returned for division by zero, modulus by zero,
	// and similar numeric-domain failures.
	KindArithmetic
	// KindFormat is returned when parsing fails (date layout, base64,
	// URL escape, etc.).
	KindFormat
)

// String returns a stable human-readable name for the kind.
func (k ErrorKind) String() string {
	switch k {
	case KindInvalidInput:
		return "invalid_input"
	case KindNotFound:
		return "not_found"
	case KindArithmetic:
		return "arithmetic"
	case KindFormat:
		return "format"
	default:
		return "unknown(" + strconv.Itoa(int(k)) + ")"
	}
}

// Error is the only error type this package returns.
//
// Callers use errors.As to extract it and switch on Kind. The four
// package-level sentinels (ErrInvalidInput, ErrNotFound, ErrArithmetic,
// ErrFormat) participate in errors.Is by Kind only — Op, Path, and Cause
// are diagnostic context, not part of identity.
type Error struct {
	Kind  ErrorKind
	Op    string // function name where the error was produced; "" allowed
	Path  string // accessor path (e.g. "users.3.name"); "" allowed
	Cause error  // underlying error, if any
}

// Error renders a stable message: "filter.<Op>: <kind>[: <path>][: <cause>]".
func (e *Error) Error() string {
	var b strings.Builder
	if e.Op != "" {
		b.WriteString(e.Op)
		b.WriteString(": ")
	}
	b.WriteString(e.Kind.String())
	if e.Path != "" {
		b.WriteString(": ")
		b.WriteString(e.Path)
	}
	if e.Cause != nil {
		b.WriteString(": ")
		b.WriteString(e.Cause.Error())
	}
	return b.String()
}

// Unwrap returns the underlying cause for use with errors.Is / errors.As.
func (e *Error) Unwrap() error { return e.Cause }

// Is reports whether the target error has the same Kind. This makes
// errors.Is(err, ErrNotFound) work on any *Error with Kind==KindNotFound,
// regardless of Op, Path, or Cause.
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.Kind == t.Kind
}

// Sentinels for kind-only matching with errors.Is.
//
// They are zero-value *Error{} with only Kind set; never inspect Op/Path
// on these.
var (
	ErrInvalidInput = &Error{Kind: KindInvalidInput}
	ErrNotFound     = &Error{Kind: KindNotFound}
	ErrArithmetic   = &Error{Kind: KindArithmetic}
	ErrFormat       = &Error{Kind: KindFormat}
)

// invalidInput builds an InvalidInput error rooted at op, optionally wrapping
// cause. cause may be nil.
func invalidInput(op string, cause error) *Error {
	return &Error{Kind: KindInvalidInput, Op: op, Cause: cause}
}

// notFound builds a NotFound error rooted at op for the given path.
func notFound(op, path string, cause error) *Error {
	return &Error{Kind: KindNotFound, Op: op, Path: path, Cause: cause}
}

// arithmetic builds an Arithmetic error rooted at op.
func arithmetic(op string, cause error) *Error {
	return &Error{Kind: KindArithmetic, Op: op, Cause: cause}
}

// formatErr builds a Format error rooted at op, wrapping cause.
func formatErr(op string, cause error) *Error {
	return &Error{Kind: KindFormat, Op: op, Cause: cause}
}

// Sentinel cause used internally to give common arithmetic failures a stable
// identity for errors.Is checks against the typed sentinel chain.
var (
	errDivisionByZero = errors.New("division by zero")
	errModulusByZero  = errors.New("modulus by zero")
)
