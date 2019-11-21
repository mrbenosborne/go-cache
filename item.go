package cache

// Item holds data relating to an item in the
// store.
type Item struct {
	data interface{}
}

// String return the value as type string.
func (i *Item) String() (string, bool) {
	ok, v := i.data.(string)
	return ok, v
}