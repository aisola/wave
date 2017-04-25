package wave

// Default is a Wave instance with FileStorage set to nil. It is intended that
// users of the default pull in the side effect of a backend package, which must
// set the backend for this instance's method.
var Default = NewWave(nil)

// AddFeature is a convenience function that calls Default.AddFeature.
func AddFeature(feature *Feature) {
	Default.AddFeature(feature)
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
func Open(info interface{}) error {
	return Default.Open(info)
}
