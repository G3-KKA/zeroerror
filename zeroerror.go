// Packgae idea is to separate static errors,
// that are used to determine behavior if error happened.
//
// From context-of-error, like dynamic data, trace and/or messages,
// that, on the other hand, used for logging and debugging.
//
// This is [zerolog] specific package,
// so it is not suitable to use with any other logger,
// or for different purpose.
//
// Compatable with [errors] package.
//
// Any mention of a "context" in this package means messages or values,
// that stored inside [ZeroError], and not [context.Context].
package zeroerr

import (
	"errors"

	"github.com/rs/zerolog"
)

const (
	// Default key for first message.
	FirstKey = "root"
	// Defaylt key for string message.
	MessageKey = "msg"
	// Default key for value message.
	ValueKey = "data"
	// Default key for error.
	ErrorKey = "error"

	// Default key when dictionary logged out.
	InsertedKey = "errcontext"
)

// ZeroError is a container for separately stored static error and logging context.
//
// Any mention of a "context" in this package means messages or values,
// that stored inside [ZeroError], and not [context.Context].
type ZeroError struct {
	err   error
	event *zerolog.Event
}

// New returns ready to use ZeroError.
//
// Fields:
//   - err filed initialized with 0 len string error, not nil.
//   - event field initialized with ready-to-use zerolog.Dict(), not nil.
//
// # It is better for maintenance and readability to use [With*] functions !
func New() *ZeroError {
	return &ZeroError{
		err:   errosempty,
		event: zerolog.Dict(),
	}
}

// # WithMsg adds msg to newly initialized ZeroError.
//
// If called on already initialized [ZeroError] then [WithMsg] just add more info.
//
// It is guaranteed:
//   - if message is first in event -- it will be with [FirstKey] key.
//   - added msg will be with [MessageKey] key.
func WithMsg(err error, msg string) *ZeroError {

	ze := New()
	if errors.As(err, ze) {
		ze.event = ze.event.Str(MessageKey, msg)

		return ze
	}

	event := zerolog.Dict().Str(FirstKey, msg)
	ze = &ZeroError{
		err:   err,
		event: event,
	}

	return ze
}

// # WithVal adds value to newly initialized ZeroError.
//
// # Be careful, only exported fields of a struct will be added !!!
//
// If called on already initialized [ZeroError] then [WithVal] just add more info.
//
// It is guaranteed:
//   - if message is first in event -- it will be with [FirstKey] key.
//   - added val will be with [ValueKey] key.
func WithVal(err error, val any) *ZeroError {

	ze := New()
	if errors.As(err, ze) {
		ze.event = ze.event.Any(ValueKey, val)

		return ze
	}

	event := zerolog.Dict().Any(FirstKey, val)
	ze = &ZeroError{
		err:   err,
		event: event,
	}

	return ze
}

// # WithKeyVal adds key-value pair to newly initialized ZeroError.
//
// If called on already initialized [ZeroError] then [WithKeyVal] just add more info.
//
// It is guaranteed:
//   - even if this is the rirst message of the event, key will be used,
//     insted of [FirstKey] key.
func WithKeyVal(err error, key string, val any) *ZeroError {

	ze := New()
	if errors.As(err, ze) {

		ze.event = ze.event.Any(key, val)

		return ze

	}

	event := zerolog.Dict().Any(key, val)
	ze = &ZeroError{
		err:   err,
		event: event,
	}

	return ze
}

// # WithKeyMsg adds key-msg pair to newly initialized ZeroError.
//
// If called on already initialized [ZeroError] then [WithKeyMsg] just add more info.
//
// It is guaranteed:
//   - even if this is the rirst message of the event, key will be used,
//     insted of [FirstKey] key.
func WithKeyMsg(err error, key, msg string) *ZeroError {

	ze := New()
	if errors.As(err, ze) {
		ze.event = ze.event.Str(key, msg)

		return ze
	}

	event := zerolog.Dict().Str(key, msg)
	ze = &ZeroError{
		err:   err,
		event: event,
	}

	return ze

}

// # TryInsert retuned func will add static error, followed by context to zerolog.Event.
//
//	Usage: logger.Debug().Func(TryInsert(err)).Send()  .
//
// It is guaranteed that key will be [InsertedKey] key.
//
// If error is not an ZeroError -- returned function is a no-op.
func TryInsert(err error) func(*zerolog.Event) {
	ze := New()
	if errors.As(err, ze) {
		return ze.Insert
	}

	return noop
}

// Insert will add static error, followed by context to zerolog.Event.
func (ze *ZeroError) Insert(e *zerolog.Event) {
	e.Err(ze.err).Dict("errcontext", ze.event)
}

// Join like errors package.
//
// Pipelining supported.
func (ze *ZeroError) Join(err error) *ZeroError {
	ze.err = errors.Join(ze.err, err)

	return ze
}
