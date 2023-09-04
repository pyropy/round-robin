package roundrobin

import (
	"errors"
	"sync/atomic"
)

// ErrServersNotExists is the error that item does not exist
var ErrItemDoesNotExist = errors.New("item does not exist")

// RoundRobin is an interface for representing round-robin balancing.
type RoundRobin[T any] interface {
	Next() T
}

type roundrobin[T any] struct {
	items []T
	next  uint32
}

// New returns RoundRobin implementation(*roundrobin).
func New[T any](items ...T) (RoundRobin[T], error) {
	if len(items) == 0 {
		return nil, ErrItemDoesNotExist
	}

	return &roundrobin[T]{
		items: items,
	}, nil
}

// Next returns next items
func (r *roundrobin[T]) Next() T {
	n := atomic.AddUint32(&r.next, 1)
	return r.items[(int(n)-1)%len(r.items)]
}
