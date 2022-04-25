package functional

// Construct a new array from those elements of values for which the
// include function returns true.
func Filter[R any](values []R, include func(v R) bool) (filtered []R) {
	for _, v := range values {
		if include(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Construct a new array from those elements of values converted to the new
// value returned by the convert function.
func Map[I any, R any](values []I, convert func(value I) R) (mapped []R) {
	mapped = make([]R, len(values))

	for i, v := range values {
		mapped[i] = convert(v)
	}

	return mapped
}

// Construct a new array from the characters in value converted to a new
// value returned by the convert function.
func MapString[R any](value string, convert func(value rune) R) (mapped []R) {
	mapped = make([]R, len(value))

	for i, v := range value {
		mapped[i] = convert(v)
	}

	return mapped
}
