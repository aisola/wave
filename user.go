package wave

// User must be implemented by any type wishing to have feature permissions
// managed by wave. For instance, your application's User must implement this.
type User interface {
	// ID returns the unique identitfier for the particular user.
	ID() string

	// Groups returns a string slice of group names to which the user belongs.
	Groups() []string
}
