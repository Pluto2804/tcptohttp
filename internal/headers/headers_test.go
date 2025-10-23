package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParser(t *testing.T) {
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\nFooFoo:  barbar\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers.Get("HOST"))
	assert.Equal(t, "barbar", headers.Get("FooFoo"))
	assert.Equal(t, "", headers.Get("MissingKey"))
	assert.Equal(t, 42, n)
	assert.True(t, done)
	//Test invalid Spacing in Headers
	headers = NewHeaders()
	data = []byte("   Host : localhost:42069    \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nHost: localhost:42069\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069,localhost:42069", headers.Get("HOST"))
	assert.False(t, done)

}
