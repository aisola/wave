package wave_test

import (
	"errors"
	"testing"

	"github.com/aisola/wave"
	"github.com/aisola/wave/wavetest"
)

func TestDefaultOpen(t *testing.T) {
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	wave.Default.SetBackend("testing")

	// test successful open
	if err := wave.Open(""); err != nil {
		t.Errorf("Default.Open unexpected error. Expected %v, got %v.", nil, err)
	}

	// test failed open
	openError := errors.New("open error")
	tb.OpenSideEffect = openError
	if err := wave.Open(""); err != openError {
		t.Errorf("Default.Open unexpected error. Expected %v, got %v.", openError, err)
	}
}

func TestDefaultClose(t *testing.T) {
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	wave.Default.SetBackend("testing")

	// test successful close
	if err := wave.Close(); err != nil {
		t.Errorf("Default.Close unexpected error. Expected %v, got %v.", nil, err)
	}

	// test failed close
	closeError := errors.New("close error")
	tb.CloseSideEffect = closeError
	if err := wave.Close(); err != closeError {
		t.Errorf("Default.Open unexpected error. Expected %v, got %v.", closeError, err)
	}
}


func TestDefaultAddFeature(t *testing.T) {
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	wave.Default.SetBackend("testing")

	name := "test"
	feature := &wave.Feature{Name: name}
	wave.AddFeature(feature)

	if err := wave.AddFeature(feature); err != nil {
		t.Errorf("Default.AddFeature unexpected error. Expected %v, got %v.", nil, err)
	}

	if tb.Features[name] != feature {
		t.Errorf("Deafault.AddFeature Expected %v, got %v", feature, tb.Features[name])
	}

	setError := errors.New("set error")
	tb.SetSideEffect = setError
	if err := wave.AddFeature(feature); err != setError {
		t.Errorf("Default.AddFeature unexpected error. Expected %v, got %v.", setError, err)
	}
}

func TestDefaultCan(t *testing.T) {
	user := newTestUser(make([]string, 0))
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	wave.Default.SetBackend("testing")

	tb.Features["test_can"] = &wave.Feature{Name: "test_can", Users: []string{user.UUID}}
	tb.Features["test_cant"] = &wave.Feature{Name: "test_cant", Users: []string{}}

	if !wave.Can(user, "test_can") {
		t.Errorf("User CANNOT access feature test_can.")
	}

	if wave.Can(user, "test_cant") {
		t.Errorf("User CAN access feature test_cant")
	}

	if wave.Can(user, "test-non-exsit") {
		t.Errorf("User CAN access non-existent feature when UndefinedAccess is false")
	}

	wave.Default.UndefinedAccess = true

	if !wave.Can(user, "test-non-exsit") {
		t.Errorf("User CANT access non-existent feature when UndefinedAccess is true")
	}
}
