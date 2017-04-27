package wave

import (
	"errors"
)

var (
	// ErrFeatureNotFound should be returned by FeatureBackend.Get if the
	// feature cannot be found.
	ErrFeatureNotFound = errors.New("feature not found")
)

// FeatureBackend is the interface which different backends must implement in
// order to be used with wave.
type FeatureBackend interface {
	// Open should take in a single argument of any type that it needs to
	// connect to its backend. Then it must connect to the backend or return
	// and error For instance, a SQLBackend might take in a MySQL connection
	// string and return connection timeout errors.
	Open(string) error

	// Close is called by the user to close the backend connection.
	Close() error

	// Get takes in a feature name string and returns the feature that
	// matches that name, or it should return ErrFeatureNotFound.
	Get(string) (*Feature, error)

	// Set is called by a Wave instance in order to add a feature to the
	// backend. The first argument is the feature name and the second is
	// the feature.
	Set(string, *Feature) error
}

// InMemoryBackend is a FeatureBackend which simply holds the feature
// information in-memory. This is just a default so that it is there, this
// really shouldn't be used in real serious software.
type InMemoryBackend struct {
	features map[string]*Feature
}

// NewInMemoryBackend creates a new InMemoryBackend.
func NewInMemoryBackend() *InMemoryBackend {
	return &InMemoryBackend{
		features: make(map[string]*Feature),
	}
}

// Close wipes and clears the in-memory feature storage.
func (imb *InMemoryBackend) Close() error {
	imb.features = make(map[string]*Feature)
	return nil
}

// Get returns a feature by name if it exists.
func (imb *InMemoryBackend) Get(name string) (*Feature, error) {
	if value, ok := imb.features[name]; ok {
		return value, nil
	}
	return nil, ErrFeatureNotFound
}

// Open literally does nothing. This is here to implement the FeatureBackend.
func (imb *InMemoryBackend) Open(string) error { return nil }

// Set will create a feature given a feature name and a wave.Feature.
func (imb *InMemoryBackend) Set(name string, feature *Feature) error {
	imb.features[name] = feature
	return nil
}
