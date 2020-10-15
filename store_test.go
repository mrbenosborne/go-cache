// +build unit

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_Add(t *testing.T) {
	s := Store{
		items: make(map[string]*Item),
	}

	cases := []struct {
		name  string
		key   string
		input interface{}
	}{
		{
			name:  "string",
			key:   "foo",
			input: "bar",
		},
		{
			name:  "int",
			key:   "foo-int",
			input: 1,
		},
		{
			name:  "bool",
			key:   "foo-bool",
			input: true,
		},
		{
			name: "struct",
			key:  "foo-struct",
			input: struct {
				name string
			}{
				name: "Foo",
			},
		},
		{
			name:  "float",
			key:   "foo-float",
			input: float32(0.55),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s.Add(tc.key, tc.input)
			s.mu.RLock()
			assert.NotNil(t, s.items[tc.key])
			assert.Equal(t, &Item{data: tc.input}, s.items[tc.key])
			s.mu.RUnlock()
		})
	}
}

func TestStore_Get(t *testing.T) {
	s := &Store{
		items: make(map[string]*Item),
	}

	s.Add("foo", "bar")
	i := s.Get("foo")
	assert.NotNil(t, i)
}

func TestStore_Delete(t *testing.T) {
	s := &Store{
		items: make(map[string]*Item),
	}

	s.Add("foo", "bar")
	i := s.Get("foo")
	assert.NotNil(t, i)

	s.Delete("foo")
	i = s.Get("foo")
	assert.Nil(t, i)
}
