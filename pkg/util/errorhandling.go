package util

// This function will return only the output value, by ignoring
// the errors. However, if one of the errors is not nill, the
// function will panic, resulting in the crash of the application.
func AssumeNoError[K any](v K, e ...error) K {
	for _, err := range e {
		if err != nil {
			panic(err)
		}
	}

	return v
}
