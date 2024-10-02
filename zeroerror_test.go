package zeroerr

import (
	"bytes"
	"errors"
	"strconv"
	"testing"

	"github.com/rs/zerolog"

	"github.com/stretchr/testify/assert"
)

var (
	_ErrStatic   = errors.New("static error")
	_TestMessage = "test__message__123"
	_TestStruct  = somestrcut{
		A: 42,
		B: _TestMessage,
	}
)

type somestrcut struct {
	A int16
	B string
}

func nonZeroErrorReturn(t *testing.T) error {
	t.Helper()
	return _ErrStatic
}
func zeroErrorReturnMsg(t *testing.T) error {
	t.Helper()
	return WithMsg(_ErrStatic, _TestMessage)
}
func zeroErrorReturnValue(t *testing.T) error {
	t.Helper()
	return WithVal(_ErrStatic, _TestStruct)
}
func TestUsageMsg(t *testing.T) {
	t.Parallel()
	buf := bytes.Buffer{}
	buf.Grow(1000)
	l := zerolog.New(&buf)
	err := zeroErrorReturnMsg(t)

	l.Debug().Func(TryInsert(err)).Send()
	assert.Contains(t, buf.String(), _TestMessage)

}
func TestUsageValue(t *testing.T) {
	t.Parallel()
	buf := bytes.Buffer{}
	buf.Grow(1000)
	l := zerolog.New(&buf)

	err := zeroErrorReturnValue(t)
	l.Debug().Func(TryInsert(err)).Send()

	assert.Contains(t, buf.String(), strconv.Itoa(int(_TestStruct.A)))
	assert.Contains(t, buf.String(), _TestStruct.B)

}
func TestAlreadyInitialized(t *testing.T) {
	t.Parallel()
	buf := bytes.Buffer{}
	buf.Grow(1000)
	l := zerolog.New(&buf)

	err := zeroErrorReturnValue(t)   // root key
	err = WithMsg(err, _TestMessage) // msg key
	err = WithVal(err, _TestStruct)  // data key
	l.Debug().Func(TryInsert(err)).Send()

	assert.Contains(t, buf.String(), _TestMessage)
	assert.Contains(t, buf.String(), strconv.Itoa(int(_TestStruct.A)))
	assert.Contains(t, buf.String(), _TestStruct.B)
	assert.Contains(t, buf.String(), FirstKey)
	assert.Contains(t, buf.String(), MessageKey)
	assert.Contains(t, buf.String(), ValueKey)
}
func TestCompatability(t *testing.T) {
	t.Parallel()

	err := zeroErrorReturnMsg(t)
	flag := errors.Is(err, _ErrStatic)
	assert.True(t, flag)
	assert.ErrorIs(t, err, _ErrStatic)

}
