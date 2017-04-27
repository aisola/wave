package wave

func init() {
	imb := NewInMemoryBackend()
	registeredBackends[""] = imb
	registeredBackends["memory"] = imb
	Default.SetBackend("")
}

// Default is a Wave instance using the default InMemoryBackend.
var Default = &Wave{}

// AddFeature is a convenience function that calls Default.AddFeature.
func AddFeature(feature *Feature) error {
	return Default.AddFeature(feature)
}

// Can is a convenience function that calls Default.Can.
func Can(user User, name string) bool {
	return Default.Can(user, name)
}

// Close is a convenience function that calls Default.Close.
func Close() error {
	return Default.Close()
}

// Open is a convenience function that calls Default.Open.
func Open(connection string) error {
	return Default.Open(connection)
}
