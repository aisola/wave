package wave

import (
	"testing"
)

func TestNewInMemoryBackend(t *testing.T) {
	var imb interface{} = NewInMemoryBackend()
	if _, ok := imb.(*InMemoryBackend); !ok {
		t.Errorf("NewInMemoryBackend did not return *InMemoryBackend")
	}
}

func TestInMemoryClose(t *testing.T) {
	imb := NewInMemoryBackend()
	if err := imb.Close(); err != nil {
		t.Errorf("InMemoryBackend.Close must always return nil")
	}
}

func TestInMemoryOpen(t *testing.T) {
	imb := NewInMemoryBackend()
	if err := imb.Open(""); err != nil {
		t.Errorf("InMemoryBackend.Open must always return nil")
	}
}

func TestInMemorySet(t *testing.T) {
	imb := NewInMemoryBackend()
	imb.features = make(map[string]*Feature)
	feature := &Feature{Name: "test"}
	imb.Set("test", feature)

	if value, ok := imb.features["test"]; !ok {
		t.Errorf("Feature didn't make it into the InMemoryBackend cache.")
	} else {
		if value != feature {
			t.Errorf("Feature does not match the one passed into InMemoryBackend.Set")
		}
	}
}

func TestInMemoryGet(t *testing.T) {
	imb := NewInMemoryBackend()
	feature := &Feature{Name: "test"}
	imb.features = make(map[string]*Feature)
	imb.features["test"] = feature

	if f, err := imb.Get("test"); err != nil {
		t.Errorf("InMemoryBackend.Get returned an error, %s", err)
	} else if f.Name != "test" {
		t.Errorf("InMemoryBackend.Get returned a magicly appearing feature")
	}

	if _, err := imb.Get("this-feature-does-not-exist"); err != ErrFeatureNotFound {
		t.Errorf("Error on InMemoryBackend.Get was %v, expecting %v", err, ErrFeatureNotFound)
	}

}
