package zeroerr

import "errors"

// ==== Compatablility functions ==== .

// Error implements error interface.
func (ze ZeroError) Error() string {
	return ze.err.Error()
}

// Unwrap -- errors package compatibility.
func (ze *ZeroError) Unwrap() error {
	return ze.err
}

// Is -- errors package compatibility.
func (ze *ZeroError) Is(err error) bool {
	return errors.Is(ze.err, err)
}

// As -- errors package compatibility.
func (ze *ZeroError) As(target any) bool {
	if ze2, ok := target.(*ZeroError); ok {
		ze2.err = ze.err
		ze2.event = ze.event

		return true
	}

	return false

}
