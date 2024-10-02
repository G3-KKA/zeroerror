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
	_ErrStatic2  = errors.New("static error2")
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
	l := zerolog.New(&buf)

	err := zeroErrorReturnMsg(t)
	l.Debug().Func(TryInsert(err)).Send()
	assert.Contains(t, buf.String(), _TestMessage)

}
func TestUsageValue(t *testing.T) {
	t.Parallel()

	buf := bytes.Buffer{}
	l := zerolog.New(&buf)

	err := zeroErrorReturnValue(t)
	l.Debug().Func(TryInsert(err)).Send()

	assert.Contains(t, buf.String(), strconv.Itoa(int(_TestStruct.A)))
	assert.Contains(t, buf.String(), _TestStruct.B)

}
func TestAlreadyInitialized(t *testing.T) {
	t.Parallel()

	buf := bytes.Buffer{}
	l := zerolog.New(&buf)

	err := zeroErrorReturnValue(t)   // root key.
	err = WithMsg(err, _TestMessage) // msg key.
	err = WithVal(err, _TestStruct)  // data key.
	l.Debug().Func(TryInsert(err)).Send()

	assert.Contains(t, buf.String(), _TestMessage)
	assert.Contains(t, buf.String(), strconv.Itoa(int(_TestStruct.A)))
	assert.Contains(t, buf.String(), _TestStruct.B)
	assert.Contains(t, buf.String(), FirstKey)
	assert.Contains(t, buf.String(), MessageKey)
	assert.Contains(t, buf.String(), ValueKey)
}
func TestAlreadyInitializedKeys(t *testing.T) {
	t.Parallel()

	const (
		keymsg = "keeey1"
		keyval = "kval"
	)

	buf := bytes.Buffer{}
	l := zerolog.New(&buf)

	err := WithKeyMsg(_ErrStatic, keymsg, _TestMessage) // msg key.
	err = WithKeyVal(err, keyval, _TestStruct)          // data key.
	l.Debug().Func(TryInsert(err)).Send()

	assert.Contains(t, buf.String(), _TestMessage)
	assert.Contains(t, buf.String(), strconv.Itoa(int(_TestStruct.A)))
	assert.Contains(t, buf.String(), _TestStruct.B)

	assert.NotContains(t, buf.String(), FirstKey)
	assert.NotContains(t, buf.String(), MessageKey)
	assert.NotContains(t, buf.String(), ValueKey)

	assert.Contains(t, buf.String(), keymsg)
	assert.Contains(t, buf.String(), keyval)
}

// Same as previous, but inverted functions.
// For coverage.
func TestAlreadyInitializedKeysInverted(t *testing.T) {
	t.Parallel()

	const (
		keymsg = "keeey1"
		keyval = "kval"
	)

	buf := bytes.Buffer{}
	l := zerolog.New(&buf)

	err := WithKeyVal(_ErrStatic, keyval, _TestStruct) // data key.
	err = WithKeyMsg(err, keymsg, _TestMessage)        // msg key.
	l.Debug().Func(TryInsert(err)).Send()

	assert.Contains(t, buf.String(), _TestMessage)
	assert.Contains(t, buf.String(), strconv.Itoa(int(_TestStruct.A)))
	assert.Contains(t, buf.String(), _TestStruct.B)

	assert.NotContains(t, buf.String(), FirstKey)
	assert.NotContains(t, buf.String(), MessageKey)
	assert.NotContains(t, buf.String(), ValueKey)

	assert.Contains(t, buf.String(), keymsg)
	assert.Contains(t, buf.String(), keyval)
}
func TestCompatability(t *testing.T) {
	t.Parallel()

	err := zeroErrorReturnMsg(t)
	flag := errors.Is(err, _ErrStatic)

	assert.True(t, flag)
	assert.ErrorIs(t, err, _ErrStatic)

}
func TestCompatabilityErrorsAs(t *testing.T) {
	t.Parallel()

	err := zeroErrorReturnMsg(t)
	ze := New()

	if errors.As(err, &ze) {
		assert.ErrorIs(t, err, _ErrStatic)
	} else {
		t.Fail()
	}

}
func TestJoin(t *testing.T) {
	var err error
	ze := WithMsg(_ErrStatic, _TestMessage).Join(_ErrStatic2)
	err = ze
	assert.ErrorIs(t, err, _ErrStatic)
	assert.ErrorIs(t, err, _ErrStatic2)

}
func TestCompatabilityErrorsAs2(t *testing.T) {
	err := _ErrStatic
	ze := WithMsg(_ErrStatic2, _TestMessage)
	if errors.As(err, ze) {
		t.Fail()
	}

}
