package wavetest

import (
	"github.com/aisola/wave"
)

// TestingBackend is a wave.FeatureBackend which is designed to be used for
// use in test cases. Features is where the added features are stored.
// CloseSideEffect, OpenSideEffect, and SetSideEffect are the errors that
// should be returned on the Close, Open, and Set method calls, respecitvely.
type TestingBackend struct {
	Features map[string]*wave.Feature
	CloseSideEffect error
	OpenSideEffect error
	SetSideEffect error
}

// NewTestBackend returns a new initialized instance of TestingBackend which is
// ready for use in test cases.
func NewTestingBackend() *TestingBackend {
	return &TestingBackend{
		Features: make(map[string]*wave.Feature),
	}
}

// Get returns a feature associated with the given name from the set of
// features. If the feature is not found the error will be
// wave.ErrFeatureNotFound.
func (tb *TestingBackend) Get(name string) (*wave.Feature, error) {
	if feature, ok := tb.Features[name]; ok {
		return feature, nil
	}
	return nil, wave.ErrFeatureNotFound
}

// Close will wipe and refresh the TestingBackend. If
// TestingBackend.CloseSideEffect is non-nil, then it will return the error.
func (tb *TestingBackend) Close() error {
	tb.Features = make(map[string]*wave.Feature)
	return tb.CloseSideEffect
}

// Open will return the TestingBackend.OpenSideEffect.
func (tb *TestingBackend) Open(connection string) error {
	// Usually a backend would connect here, but we don't have anything
	// to connect to.
	return tb.OpenSideEffect
}

// Set will set the feature and associated name if TestingBackend.SetSideEffect
// is nil, otherwise the feature is not set and SetSideEffect is returned.
func (tb *TestingBackend) Set(name string, feature *wave.Feature) error {
	if tb.SetSideEffect != nil {
		// If there's an error, we don't want to set it, so return early.
		return tb.SetSideEffect
	}

	tb.Features[name] = feature
	return nil
}
