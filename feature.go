package wave

// Feature is a specific feature and its constraints.
type Feature struct {
	Name   string
	Groups []string
	Users  []string
}

// NewFeatureGroups will create a new feature name, which only allows access to
// users belonging to the listed groups.
func NewFeatureGroups(name string, groups []string) *Feature {
	return &Feature{
		Name:   name,
		Groups: groups,
		Users:  nil,
	}
}

// NewFeatureUsers will create a new feature name, which only allows access to
// the provided users.
func NewFeatureUsers(name string, users []string) *Feature {
	return &Feature{
		Name:   name,
		Groups: nil,
		Users:  users,
	}
}

// Can returns true if the user has access to this feature based on the
// contraints. This method checks if the user is allowed specificly by ID
// first, then by group.
func (f *Feature) Can(user User) bool {
	if f.canUserID(user) {
		return true
	} else if f.canGroup(user) {
		return true
	}
	return false
}

func (f *Feature) canGroup(user User) bool {
	if f.Groups == nil || stringInSlice("NONE", f.Groups) {
		return false
	}

	if stringInSlice("ALL", f.Groups) {
		return true
	}

	userGroups := user.Groups()

	for _, group := range f.Groups {
		if stringInSlice(group, userGroups) {
			return true
		}
	}

	return false
}

func (f *Feature) canUserID(user User) bool {
	if f.Users == nil {
		return false
	}

	if stringInSlice(user.ID(), f.Users) {
		return true
	}

	return false
}
