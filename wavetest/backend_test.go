package wavetest

import (
	"errors"
	"testing"

	"github.com/aisola/wave"
)

func TestTestingBackend(t *testing.T) {
	var tb interface{} = NewTestingBackend()
	if _, ok := tb.(*TestingBackend); !ok {
		t.Errorf("NewTestingBackend did not return *TestingBackend")
	}
}

func TestTestingBackendClose(t *testing.T) {
	tb := NewTestingBackend()

	if err := tb.Close(); err != nil {
		t.Errorf("TestingBackend.Close unexpected error. Expected %v, got %v.", nil, err)
	}

	closeError := errors.New("close error")
	tb.CloseSideEffect = closeError

	if err := tb.Close(); err != closeError {
		t.Errorf("TestingBackend.Close unexpected error. Expected %v, got %v.", closeError, err)
	}
}

func TestTestingBackendOpen(t *testing.T) {
	tb := NewTestingBackend()

	if err := tb.Open(""); err != nil {
		t.Errorf("TestingBackend.Open unexpected error. Expected %v, got %v.", nil, err)
	}

	openError := errors.New("open error")
	tb.OpenSideEffect = openError

	if err := tb.Open(""); err != openError {
		t.Errorf("TestingBackend.Open unexpected error. Expected %v, got %v.", openError, err)
	}
}

func TestTestingBackendSet(t *testing.T) {
	tb := NewTestingBackend()
	feature := &wave.Feature{Name: "test"}

	if err := tb.Set("test", feature); err != nil {
		t.Errorf("TestingBackend.Set unexpected error. Expected %v, got %v.", nil, err)
	}

	if value, ok := tb.Features["test"]; !ok {
		t.Errorf("TestingBackend.Set succeeded, but feature not found.")
	} else {
		if value != feature {
			t.Errorf("TestingBackend.Set succeeded, but feature in backend does not match the one that was set.")
		}
	}

	setError := errors.New("set error")
	tb.Features = make(map[string]*wave.Feature) // reset the backend
	tb.SetSideEffect = setError

	if err := tb.Set("test", feature); err != setError {
		t.Errorf("TestingBackend.Set unexpected error. Expected %v, got %v.", setError, err)
	}

	if _, ok := tb.Features["test"]; ok {
		t.Errorf("TestingBackend.Set failed, but feature still set.")
	}
}

func TestTestingBackendGet(t *testing.T) {
	tb := NewTestingBackend()
	feature := &wave.Feature{Name: "test"}
	tb.Features["test"] = feature

	if f, err := tb.Get("test"); err != nil {
		t.Errorf("TestingBackend.Get unexpected error. Expected %v, got %v.", nil, err)
	} else if f.Name != "test" {
		t.Errorf("TestingBackend.Get succeeded, but returned a magicly appearing feature")
	}

	if _, err := tb.Get("this-feature-does-not-exist"); err != wave.ErrFeatureNotFound {
		t.Errorf("TestingBackend.Get unexpected error. Expected %v, got %v.", wave.ErrFeatureNotFound, err)
	}
}
