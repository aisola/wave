package memory

import (
	"testing"

	"github.com/aisola/wave"
)

func TestNewInMemoryBackend(t *testing.T) {
	if _, ok := NewInMemoryBackend().(*InMemoryBackend); !ok {
		t.Errorf("NewInMemoryBackend did not return *InMemoryBackend")
	}
}

func TestClose(t *testing.T) {
	imb := NewInMemoryBackend()
	if err := imb.Close(); err != nil {
		t.Errorf("InMemoryBackend.Close must always return nil")
	}
}

func TestOpen(t *testing.T) {
	imb := NewInMemoryBackend()
	if err := imb.Open(nil); err != nil {
		t.Errorf("InMemoryBackend.Open must always return nil")
	}
}

func TestSet(t *testing.T) {
	imb := NewInMemoryBackend().(*InMemoryBackend)
	imb.features = make(map[string]*wave.Feature)
	feature := &wave.Feature{Name: "test"}
	imb.Set("test", feature)

	if value, ok := imb.features["test"]; !ok {
		t.Errorf("Feature didn't make it into the InMemoryBackend cache.")
	} else {
		if value != feature {
			t.Errorf("Feature does not match the one passed into InMemoryBackend.Set")
		}
	}
}

func TestGet(t *testing.T) {
	imb := NewInMemoryBackend().(*InMemoryBackend)
	feature := &wave.Feature{Name: "test"}
	imb.features = make(map[string]*wave.Feature)
	imb.features["test"] = feature

	if f, err := imb.Get("test"); err != nil {
		t.Errorf("InMemoryBackend.Get returned an error, %s", err)
	} else if f.Name != "test" {
		t.Errorf("InMemoryBackend.Get returned a magicly appearing feature")
	}

	if _, err := imb.Get("this-feature-does-not-exist"); err != wave.ErrFeatureNotFound {
		t.Errorf("Error on InMemoryBackend.Get was %v, expecting %v", err, wave.ErrFeatureNotFound)
	}

}
