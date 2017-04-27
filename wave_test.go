package wave_test

import (
	"errors"
	"testing"

	"github.com/aisola/wave"
	"github.com/aisola/wave/wavetest"
)

func TestWaveOpen(t *testing.T) {
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	features := wave.New("testing")

	// test successful open
	if err := features.Open(""); err != nil {
		t.Errorf("Wave.Open unexpected error. Expected %v, got %v.", nil, err)
	}

	// test failed open
	openError := errors.New("open error")
	tb.OpenSideEffect = openError
	if err := features.Open(""); err != openError {
		t.Errorf("Wave.Open unexpected error. Expected %v, got %v.", openError, err)
	}
}

func TestWaveClose(t *testing.T) {
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	features := wave.New("testing")

	// test successful close
	if err := features.Close(); err != nil {
		t.Errorf("Wave.Close unexpected error. Expected %v, got %v.", nil, err)
	}

	closeError := errors.New("close error")
	tb.CloseSideEffect = closeError
	// test failed close
	if err := features.Close(); err != closeError {
		t.Errorf("Wave.Close unexpected error. Expected %v, got %v.", closeError, err)
	}
}

func TestWaveAddFeature(t *testing.T) {
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	features := wave.New("testing")

	name := "test"
	feature := &wave.Feature{Name: name}

	if err := features.AddFeature(feature); err != nil {
		t.Errorf("Wave.AddFeature unexpected error. Expected %v, got %v.", nil, err)
	}

	if tb.Features[name] != feature {
		t.Errorf("Wave.AddFeature Expected %v, got %v", feature, tb.Features[name])
	}

	setError := errors.New("set error")
	tb.SetSideEffect = setError
	if err := features.AddFeature(feature); err != setError {
		t.Errorf("Wave.AddFeature unexpected error. Expected %v, got %v.", setError, err)
	}
}

func TestWaveCan(t *testing.T) {
	user := newTestUser(make([]string, 0))
	tb := wavetest.NewTestingBackend()
	wave.Register("testing", tb)
	tb.Features["test_can"] = &wave.Feature{Name: "test_can", Users: []string{user.UUID}}
	tb.Features["test_cant"] = &wave.Feature{Name: "test_cant", Users: []string{}}

	features := wave.New("testing")

	if !features.Can(user, "test_can") {
		t.Errorf("User CANNOT access feature test_can.")
	}

	if features.Can(user, "test_cant") {
		t.Errorf("User CAN access feature test_cant")
	}

	if features.Can(user, "test-non-exsit") {
		t.Errorf("User CAN access non-existent feature when UndefinedAccess is false")
	}

	features.UndefinedAccess = true

	if !features.Can(user, "test-non-exsit") {
		t.Errorf("User CANT access non-existent feature when UndefinedAccess is true")
	}
}


func TestWaveSetStorage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("SetBackend didn't work, expected panic. The call didn't panic.")
		}
	}()
	tb1 := wavetest.NewTestingBackend()
	wave.Register("testing1", tb1)
	features := wave.New("testing")

	features.SetBackend("this-storage-does-not-exist")
}
