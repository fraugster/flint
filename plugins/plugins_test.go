package plugins

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dummy struct {
	counter int
	err     bool
}

func (d *dummy) Process(context.Context, []byte) error {
	d.counter++
	if d.err {
		return errors.New("random err")
	}
	return nil
}

func TestRegister(t *testing.T) {
	pl := &dummy{}
	require.NotPanics(t, func() { Register("d1", pl) })
	require.Panics(t, func() { Register("d1", pl) })

	require.NotPanics(t, func() { Register("d2", pl) })
	assert.NoError(t, Call(context.Background(), []byte("not important")))
	assert.Equal(t, 2, pl.counter)
	pl.err = true
	pl.counter = 0
	assert.Error(t, Call(context.Background(), []byte("not important")))
	assert.Equal(t, 1, pl.counter)
}
