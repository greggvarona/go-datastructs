package sets

import "github.com/greggvarona/go-datastructs/model"

// Hashable allows implementing structures to define their own implementation.
type Hashable interface {
	HashCode() string
}

// HashSet is a simple implementation of a hashed key set.
// It's recommended that the items stored here are Hashable.
type HashSet struct {
	data map[string]model.Void
}

// Put places d into the set.
func (hs *HashSet) Put(d Hashable) {
	hs.data[d.HashCode()] = model.Void{}
}

// Contains d if d exists in the set.
func (hs *HashSet) Contains(d Hashable) bool {
	_, ok := hs.data[d.HashCode()]
	return ok
}

// Delete will remove d from the set.
func (hs *HashSet) Delete(d Hashable) {
	if hs.Contains(d) {
		delete(hs.data, d.HashCode())
	}
}

// Size will return the number of elements in the set.
func (hs *HashSet) Size() int {
	return len(hs.data)
}
