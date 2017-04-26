package wave_test

import (
	"testing"

	"github.com/aisola/wave"
)

func TestNewFeatureGroups(t *testing.T) {
	featureA := &wave.Feature{Name: "test", Groups: []string{"test"}, Users: nil}
	featureB := wave.NewFeatureGroups("test", []string{"test"})

	if featureA.Name != featureB.Name {
		t.Errorf("Feature names do not match")
	}

	if featureA.Groups[0] != featureB.Groups[0] {
		t.Errorf("Feature groups do not match")
	}

	if featureB.Users != nil {
		t.Errorf("Feature users not nil")
	}
}

func TestNewFeatureUsers(t *testing.T) {
	featureA := &wave.Feature{Name: "test", Groups: nil, Users: []string{"test"}}
	featureB := wave.NewFeatureUsers("test", []string{"test"})

	if featureA.Name != featureB.Name {
		t.Errorf("Feature names do not match")
	}

	if featureA.Users[0] != featureB.Users[0] {
		t.Errorf("Feature users do not match")
	}

	if featureB.Groups != nil {
		t.Errorf("Feature groups not nil")
	}
}

func TestFeatureCan(t *testing.T) {
	user1 := newTestUser(make([]string, 0))
	user2 := newTestUser([]string{"test"})
	featureUsersNil := &wave.Feature{Name: "feature_users_nil", Groups: nil, Users: nil}
	featureUsers := &wave.Feature{Name: "feature_users", Groups: nil, Users: []string{user1.UUID}}
	featureGroupsNil := &wave.Feature{Name: "feature_groups_nil", Groups: nil, Users: nil}
	featureGroupsNone := &wave.Feature{Name: "feature_groups_none", Groups: wave.NONE, Users: nil}
	featureGroupsAll := &wave.Feature{Name: "feature_groups_all", Groups: wave.ALL, Users: nil}
	featureGroupsTest := &wave.Feature{Name: "feature_groups_test", Groups: []string{"test"}, Users: nil}

	if featureUsersNil.Can(user1) || featureUsersNil.Can(user2) {
		t.Errorf("Users CAN access feature_users_nil")
	}

	if !featureUsers.Can(user1) {
		t.Errorf("User1 CANNOT access feature_users")
	}

	if featureUsers.Can(user2) {
		t.Errorf("User2 CAN access feature_users")
	}

	if featureGroupsNil.Can(user1) || featureGroupsNil.Can(user2) {
		t.Errorf("Users CAN access feature_groups_nil")
	}

	if featureGroupsNone.Can(user1) || featureGroupsNone.Can(user2) {
		t.Errorf("Users CAN access feature_groups_none")
	}

	if !featureGroupsAll.Can(user1) || !featureGroupsAll.Can(user2) {
		t.Errorf("Users CANNOT access feature_groups_all")
	}

	if !featureGroupsTest.Can(user2) {
		t.Errorf("User2 CANNOT access feature_groups_test")
	}

	if featureGroupsTest.Can(user1) {
		t.Errorf("User1 CAN access feature_groups_test")
	}

}
