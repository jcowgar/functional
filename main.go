// Package functional provides some basic functional language
// type functions.
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

// Produce a map from an array of keys and values.
//
//     keys := []string{"name", "age"}
//     values := []any{"John", 28}
//     mapped := Zip(keys, values)
//
//     // { "name": "John", "age": 28 }
func Zip[K comparable, V any](keys []K, values []V) map[K]V {
	result := make(map[K]V)

	for i := 0; i < len(keys); i++ {
		result[keys[i]] = values[i]
	}

	return result
}

// Produce an array of keys and values from a map.
//
//     m := map[string]any{"name":"John", "age":28}
//     keys, values := Unzip(m)
//
//     // keys = { "name", "age" }
//     // values = { "John", 28 }
func Unzip[K comparable, V any](m map[K]V) (keys []K, values []V) {
	keys = make([]K, 0)
	values = make([]V, 0)

	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}

	return keys, values
}
