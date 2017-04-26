package wave_test

import (
	"testing"

	"github.com/aisola/wave"
)

func TestDefaultOpen(t *testing.T) {
	tb := newTestingBackend()
	wave.Default.SetStorage(tb)

	// test successful open
	if err := wave.Open(nil); err != nil {
		t.Errorf("Error on wave.Open was %v, expecting nil", err)
	}

	// test failed open
	if err := wave.Open(wave.ErrFeatureNotFound); err != wave.ErrFeatureNotFound {
		t.Errorf("Error on wave.Open was %v, expecting %v", err, wave.ErrFeatureNotFound)
	}
}

func TestDefaultClose(t *testing.T) {
	tb := newTestingBackend()
	wave.Default.SetStorage(tb)

	// test successful close
	if err := wave.Close(); err != nil {
		t.Errorf("Error on Close was %v, expecting nil", err)
	}

	tb.CloseSideEffect = wave.ErrFeatureNotFound
	// test failed close
	if err := wave.Close(); err != wave.ErrFeatureNotFound {
		t.Errorf("Error on Open was %v, expecting %v", err, wave.ErrFeatureNotFound)
	}
}


func TestDefaultAddFeature(t *testing.T) {
	tb := newTestingBackend()
	wave.Default.SetStorage(tb)

	name := "test"
	feature := &wave.Feature{Name: name}
	wave.AddFeature(feature)

	if tb.Features[name] != feature {
		t.Errorf("Feature test, not added correctly. Expected %v, got %v", feature, tb.Features[name])
	}
}

func TestDefaultCan(t *testing.T) {
	user := newTestUser(make([]string, 0))
	tb := newTestingBackend()
	tb.Features["test_can"] = &wave.Feature{Name: "test_can", Users: []string{user.UUID}}
	tb.Features["test_cant"] = &wave.Feature{Name: "test_cant", Users: []string{}}
	wave.Default.SetStorage(tb)

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
