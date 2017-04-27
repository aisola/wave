package wave

import (
	"fmt"
)

var registeredBackends = map[string]FeatureBackend{}

// Register will register a FeatureBackend wave.
func Register(name string, backend FeatureBackend) {
	registeredBackends[name] = backend
}

// Wave manages all of the interactions between you and the FeatureBackend.
type Wave struct {
	backend string
	storage FeatureBackend

	// When UndefinedAccess is false, the users will NOT be granted access
	// to features of undefined names.
	UndefinedAccess bool
}

// New creates a new *Wave instance and sets the given FeatureBackend.
func New(backend string) *Wave {
	wave := &Wave{UndefinedAccess: false}
	wave.SetBackend(backend)
	return wave
}

// AddFeature adds a Feature to the wave instance.
func (w *Wave) AddFeature(feature *Feature) error {
	return w.storage.Set(feature.Name, feature)
}

// Can returns true of the given user has access to the given feature.
func (w *Wave) Can(user User, name string) bool {
	feature, err := w.storage.Get(name)
	if err != nil {
		return w.UndefinedAccess
	}
	return feature.Can(user)
}

// Close will safely close up (and persist if necessary) the FeatureBackend.
func (w *Wave) Close() error {
	return w.storage.Close()
}

// Open connects to the FeatureBackend using the given connection string.
func (w *Wave) Open(connection string) error {
	return w.storage.Open(connection)
}

// SetBackend sets the FeatureBackend for this wave instance. It will panic
// if the named backend is not registered.
func (w *Wave) SetBackend(backend string) {
	storage, ok := registeredBackends[backend]
	if !ok {
		panic(fmt.Sprintf("FeatureBackend `%s` is not registered.", backend))
	}

	w.backend = backend
	w.storage = storage
}
