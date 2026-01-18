package gx

// Must panics if err is not nil, otherwise returns the value.
// This is useful for wrapping functions that return (T, error) in contexts
// where errors should be fatal, such as during initialization.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
