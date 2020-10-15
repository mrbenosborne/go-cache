package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItem_String(t *testing.T) {
	i := &Item{
		data: "foo",
	}

	s, ok := i.String()
	assert.True(t, ok)
	assert.Equal(t, "foo", s)
}
