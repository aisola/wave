package memory

import (
	"github.com/aisola/wave"
)

func init() {
	wave.Default.SetStorage(NewInMemoryBackend())
}

// InMemoryBackend is a wave.FeatureBackend which simply holds the feature
// information in-memory. This is just a default so that it is there, this
// really shouldn't be used in real serious software.
type InMemoryBackend struct {
	features map[string]*wave.Feature
}

// NewInMemoryBackend creates a new InMemoryBackend.
func NewInMemoryBackend() wave.FeatureBackend {
	return new(InMemoryBackend)
}

// Close closes the backend and desroys the features.
func (imb *InMemoryBackend) Close() error {
	imb.features = nil
	return nil
}

// Get returns a feature by name if it exists.
func (imb *InMemoryBackend) Get(name string) (*wave.Feature, error) {
	if value, ok := imb.features[name]; ok {
		return value, nil
	}
	return nil, wave.FeatureNotFoundError
}

// Open sets up the backend to be used.
func (imb *InMemoryBackend) Open(interface{}) error {
	imb.features = make(map[string]*wave.Feature)
	return nil
}

// Set will create a feature given a feature name and a wave.Feature.
func (imb *InMemoryBackend) Set(name string, feature *wave.Feature) {
	imb.features[name] = feature
}
