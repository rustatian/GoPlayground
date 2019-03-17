package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {

}

// Coding Task: Concurrent in-memory cache.
//
// Fetcher (see below) is an interface which abstracts the process of fetching
// and loading a "Model".  In practice there would be Fetcher implementations
// for retrieving and loading models from local file systems, S3 buckets etc...
//
// Implement and test an in-memory cache which wraps a given Fetcher and caches
// calls to its Fetch method (complete the implementation of NewCache and the
// FetchCache type below).

// Model is a resource.
type Model struct{}

// Fetcher is an interface that defines the Fetch method.
type Fetcher interface {
	// Fetch retrieves an Model for a given identifier id.
	Fetch(ctx context.Context, id string) (*Model, error)
}

// NewCache creates a new Fetcher which caches calls to f.Fetch.
// See FetchCache for more details.
func NewCache(f Fetcher) Fetcher {
	return &mitm{
		ff:     f,
		cache: map[string]*Model{},
	}
}

type mitm struct {
	ff     Fetcher
	cache map[string]*Model
}

func (m *mitm) Fetch(ctx context.Context, id string) (*Model, error) {
	mm, err := m.ff.Fetch(ctx, id)
	return mm, err
}

// FetchCache implements an in-memory cache for a Fetcher.
//
// A FetchCache is safe for use by multiple goroutines simultaneously.
type FetchCache struct {
	sm sync.Map
}

// Fetch implements Fetcher.
func (f *FetchCache) Fetch(ctx context.Context, id string) (*Model, error) {
	m, ok := f.sm.Load(id)
	if !ok {
		return nil, fmt.Errorf("no such key") // --> or return default value
	}

	model, ok := m.(*Model)
	if !ok {
		return nil, fmt.Errorf("type assertion to *Model error")
	}

	return model, nil
}

// Clear the cache for the given id.
func (f *FetchCache) Clear(id string) {
	f.sm.Delete(id)
}
