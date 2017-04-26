package wave

import (
	"errors"
)

var (
	// FeatureNotFoundError should be returned by a FeatureBacked's Get
	// method if the feature cannot be found.
	ErrFeatureNotFound = errors.New("feature not found")
)

// FeatureBackend is the interface which different backends must implement in
// order to be used with this package.
type FeatureBackend interface {
	// Open should take in a single argument of any type that it needs to
	// connect to its backend. Then it must connect to the backend or return
	// and error For instance, a SQLBackend might take in a MySQL connection
	// string and return connection timeout errors.
	Open(interface{}) error

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
