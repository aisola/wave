package wave_test

import (
	"fmt"
	"strconv"
)

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
