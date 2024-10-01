package v1

import (
	"errors"

	"github.com/rs/zerolog"
)

type merror struct {
	err  error
	dict *zerolog.Event
}

// Error implements error
func (m *merror) Error() string {
	return m.err.Error()
}

// IsMerror check that you can add context to the error.
func IsMerror(err error) (*zerolog.Event, bool) {
	if mer, ok := err.(*merror); ok {
		return mer.dict, true
	}
	return nil, false
}

// ToLogger for finaly printing to zerologger.
func (m *merror) ToLogger() *zerolog.Event {
	return m.dict
}

func (m *merror) AddInfo(key, msg string) {
	m.dict = m.dict.Str(key, msg)

}

// Unwrap -- errors package compatability
func (m *merror) Unwrap() error {
	return m.err
}

// Is -- errors package compatability
func (m *merror) Is(err error) bool {
	return m.err == err
}

// Join like errors package.
func (m *merror) Join(err error) error {
	m.err = errors.Join(m.err, err)
	return m.err
}
