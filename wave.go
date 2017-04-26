package wave

// Wave manages all of the interactions between you and the FeatureBackend.
type Wave struct {
	storage FeatureBackend

	// When UndefinedAccess is false, the users will NOT be granted access
	// to features of undefined names.
	UndefinedAccess bool
}

// NewWave creates a new Wave instance from the provided FeatureBackend.
func NewWave(backend FeatureBackend) *Wave {
	return &Wave{
		storage:         backend,
		UndefinedAccess: false,
	}
}

// AddFeature adds a Feature to the wave instance.
func (r *Wave) AddFeature(feature *Feature) error {
	return r.storage.Set(feature.Name, feature)
}

// Can returns true of the given user has access to the given feature.
func (r *Wave) Can(user User, name string) bool {
	feature, err := r.storage.Get(name)
	if err != nil {
		return r.UndefinedAccess
	}
	return feature.Can(user)
}

// Close will safely close up (and persist if necessary) the FeatureBackend.
func (r *Wave) Close() error {
	return r.storage.Close()
}

// Open will open (and load in if neccessary) the FeatureBackend.
func (r *Wave) Open(info interface{}) error {
	return r.storage.Open(info)
}

// SetStorage will set the storage backend. This method is useful to users for
// creating backends and configuring them after the Wave instance is created.
func (r *Wave) SetStorage(backend FeatureBackend) {
	r.storage = backend
}
