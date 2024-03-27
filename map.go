// Package gsync Description: A generic wrapper around sync.Map.
package gsync

import (
	"fmt"
	"strings"
	"sync"
)

// Map is a generic wrapper around sync.Map.
type Map[K comparable, V any] struct {
	// data is the underlying sync.Map.
	data sync.Map
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.data.Load(key)
	if !ok {
		return value, ok
	}
	return v.(V), ok
}

// Get returns Load result ignoring ok value.
func (m *Map[K, V]) Get(key K) (value V) {
	value, _ = m.Load(key)
	return
}

// Has returns true if key exists in the map.
func (m *Map[K, V]) Has(key K) bool {
	_, ok := m.data.Load(key)
	return ok
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.data.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}
	return v.(V), loaded
}

// Store sets the value for a key.
func (m *Map[K, V]) Store(key K, value V) {
	m.data.Store(key, value)
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.data.LoadOrStore(key, value)
	return a.(V), loaded
}

// GetOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map[K, V]) GetOrStore(key K, value V) (actual V) {
	a, _ := m.data.LoadOrStore(key, value)
	return a.(V)
}

// Delete deletes the value for a key.
func (m *Map[K, V]) Delete(key K) {
	m.data.Delete(key)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *Map[K, V]) Swap(key K, new V) (old V, loaded bool) {
	a, loaded := m.data.Swap(key, new)
	return a.(V), loaded
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The old value must be of a comparable type.
func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.data.CompareAndSwap(key, old, new)
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
func (m *Map[K, V]) CompareAndDelete(key K, value V) bool {
	return m.data.CompareAndDelete(key, value)
}

// Compute LoadOrCompute returns the existing value for the key if present.
func (m *Map[K, V]) Compute(key K, f func(old V) V) {
	m.Store(key, f(m.Get(key)))
}

// ComputeAndGet LoadOrCompute returns the existing value for the key if present.
func (m *Map[K, V]) ComputeAndGet(key K, f func(old V) V) V {
	return m.GetOrStore(key, f(m.Get(key)))
}

// ComputeAndLoad LoadOrCompute returns the existing value for the key if present.
func (m *Map[K, V]) ComputeAndLoad(key K, f func(old V) V) (V, bool) {
	return m.LoadOrStore(key, f(m.Get(key)))
}

// Len returns the number of items in the map.
func (m *Map[K, V]) Len() int {
	l := 0
	m.data.Range(func(key, value any) bool {
		l++
		return true
	})
	return l
}

// Keys returns all keys in the map.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0)
	m.data.Range(func(key, value any) bool {
		keys = append(keys, key.(K))
		return true
	})
	return keys
}

// Values returns all values in the map.
func (m *Map[K, V]) Values() []V {
	values := make([]V, 0)
	m.data.Range(func(key, value any) bool {
		values = append(values, value.(V))
		return true
	})
	return values
}

// Clear removes all items from the map.
func (m *Map[K, V]) Clear() {
	m.data.Range(func(key, value any) bool {
		m.data.Delete(key)
		return true
	})
}

// String returns a string representation of the map.
func (m *Map[K, V]) String() string {
	var toPrint []string
	m.Range(func(key K, value V) bool {
		toPrint = append(toPrint, fmt.Sprintf("%v: %v", key, value))
		return true
	})
	return fmt.Sprintf("{%v}", strings.Join(toPrint, ", "))
}
