package zeroerr

import (
	"errors"

	"github.com/rs/zerolog"
)

var errosempty = errors.New("")

func noop(e *zerolog.Event) {}
