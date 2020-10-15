// +build unit

package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCache_New(t *testing.T) {
	cases := []struct {
		name string
		size int
		want *Cache
	}{
		{
			name: "simple",
			size: 0,
			want: &Cache{
				Stores: map[string]*Store{},
			},
		},
		{
			name: "size 10",
			size: 10,
			want: &Cache{
				Stores: map[string]*Store{},
			},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := New(tt.size)
			assert.Equal(t, tt.want, c)
		})
	}
}

func TestCache_NewStore(t *testing.T) {
	cases := []struct {
		name  string
		input string
		size  int
		want  *Store
	}{
		{
			name:  "simple",
			input: "foo",
			size:  0,
			want: &Store{
				items: make(map[string]*Item),
			},
		},
		{
			name:  "simple",
			input: "foo and bar",
			size:  10,
			want: &Store{
				items: make(map[string]*Item),
			},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := New(tt.size)
			s := c.New(tt.input)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestCache_GetStore(t *testing.T) {
	cases := []struct {
		name  string
		input string
		size  int
		want  *Store
	}{
		{
			name:  "simple",
			input: "foo",
			size:  0,
			want: &Store{
				items: make(map[string]*Item),
			},
		},
		{
			name:  "simple",
			input: "foo and bar",
			size:  0,
			want: &Store{
				items: make(map[string]*Item),
			},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := New(tt.size)
			c.New(tt.input)
			s := c.Get(tt.input)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestCache_DeleteStore(t *testing.T) {
	c := New(0)
	c.New("foo")
	require.NotNil(t, c.Get("foo"))

	c.Delete("foo")
	assert.Nil(t, c.Get("foo"))
}
