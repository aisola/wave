package wave_test

import (
	"fmt"
	"strconv"

	"github.com/aisola/wave"
)

// FEATURE BACKEND FOR TESTING
type testingBackend struct {
	Features map[string]*wave.Feature
	CloseSideEffect error
	SetSideEffect error
}

func newTestingBackend() *testingBackend {
	return &testingBackend{Features: make(map[string]*wave.Feature)}
}

func (tb *testingBackend) Open(info interface{}) error {
	var err error
	if info == nil {
		return err
	}
	return info.(error)
}

func (tb *testingBackend) Close() error {
	return tb.CloseSideEffect
}

func (tb *testingBackend) Get(name string) (*wave.Feature, error) {
	if value, ok := tb.Features[name]; ok {
		return value, nil
	}
	return nil, wave.ErrFeatureNotFound
}

func (tb *testingBackend) Set(name string, feature *wave.Feature) error {
	tb.Features[name] = feature
	return tb.SetSideEffect
}

// USER INFORMATION FOR TESTING
var userCount = 0

type testUser struct {
	UUID string
	Name string
	FG   []string
}

func newTestUser(groups []string) *testUser {
	uuid := strconv.Itoa(userCount)
	userCount++
	return &testUser{
		UUID: uuid,
		Name: fmt.Sprintf("Test User {%s}", uuid),
		FG: groups,
	}
}

func (tu *testUser) ID() string {
	return tu.UUID
}

func (tu *testUser) Groups() []string {
	return tu.FG
}
