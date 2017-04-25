package memory

import (
	"github.com/aisola/wave"
)

func init() {
	wave.Default.SetStorage(NewInMemoryBackend())
}

type InMemoryBackend struct {
	features map[string]*wave.Feature
}

func NewInMemoryBackend() wave.FeatureBackend {
	return new(InMemoryBackend)
}

func (imb *InMemoryBackend) Close() error {
	imb.features = nil
	return nil
}

func (imb *InMemoryBackend) Get(name string) (*wave.Feature, error) {
	if value, ok := imb.features[name]; ok {
		return value, nil
	}
	return nil, wave.FeatureNotFoundError
}

func (imb *InMemoryBackend) Open(interface{}) error {
	imb.features = make(map[string]*wave.Feature)
	return nil
}

func (imb *InMemoryBackend) Set(name string, feature *wave.Feature) {
	imb.features[name] = feature
}
