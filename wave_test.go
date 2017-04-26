package wave_test

import (
	"testing"

	"github.com/aisola/wave"
)


// TestWaveInstantiation tests that the Wave struct receives the backend correctly.
// Since the Wave struct has the storage field being private, we do this by setting
// the CloseSideEffect on our testing backend, we KNOW that this is the only backend
// in this package that does that so we can consider it passed.
func TestWaveInstantiation(t *testing.T) {
	// Redo this one so that it uses just the testBackend
	tb := newTestingBackend()
	tb.CloseSideEffect = wave.ErrFeatureNotFound
	features := wave.NewWave(tb)
	if err := features.Close(); err != wave.ErrFeatureNotFound {
		t.Errorf("Error on Close was %v, expecting %v", err, wave.ErrFeatureNotFound)
	}
}

func TestWaveOpen(t *testing.T) {
	tb := newTestingBackend()
	features := wave.NewWave(tb)

	// test successful open
	if err := features.Open(nil); err != nil {
		t.Errorf("Error on Open was %v, expecting nil", err)
	}

	// test failed open
	if err := features.Open(wave.ErrFeatureNotFound); err != wave.ErrFeatureNotFound {
		t.Errorf("Error on Open was %v, expecting %v", err, wave.ErrFeatureNotFound)
	}
}

func TestWaveClose(t *testing.T) {
	tb := newTestingBackend()
	features := wave.NewWave(tb)

	// test successful close
	if err := features.Close(); err != nil {
		t.Errorf("Error on Close was %v, expecting nil", err)
	}

	tb.CloseSideEffect = wave.ErrFeatureNotFound
	// test failed close
	if err := features.Close(); err != wave.ErrFeatureNotFound {
		t.Errorf("Error on Open was %v, expecting %v", err, wave.ErrFeatureNotFound)
	}
}

func TestWaveAddFeature(t *testing.T) {
	tb := newTestingBackend()
	features := wave.NewWave(tb)

	name := "test"
	feature := &wave.Feature{Name: name}
	features.AddFeature(feature)

	if tb.Features[name] != feature {
		t.Errorf("Feature test, not added correctly. Expected %v, got %v", feature, tb.Features[name])
	}
}

func TestWaveSetStorage(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("SetStorage didn't work, expected panic. The call didn't panic.")
		}
	}()
	tb1 := newTestingBackend()
	tb2 := newTestingBackend()
	tb2.Features = nil

	features := wave.NewWave(tb1)

	features.SetStorage(tb2)

	features.AddFeature(&wave.Feature{Name: "test"})
}

func TestWaveCan(t *testing.T) {
	user := newTestUser(make([]string, 0))
	tb := newTestingBackend()
	tb.Features["test_can"] = &wave.Feature{Name: "test_can", Users: []string{user.UUID}}
	tb.Features["test_cant"] = &wave.Feature{Name: "test_cant", Users: []string{}}

	features := wave.NewWave(tb)

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
