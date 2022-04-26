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

// Iterate calling reducer on the previously calculated value and the current
// value from the array. On the first call, initial is used as the "previously
// calculated value."
//
//     nums := []int{1,2,3}
//     sum := func (previous int, current int) int { return previous + current }
//     result := Reduce(nums, sum, 0) // 6
func Reduce[V comparable](ary []V, reducer func(previous V, current V) V, initial V) V {
	result := initial

	for _, v := range ary {
		result = reducer(result, v)
	}

	return result
}

// Iterate backward calling reducer on the previously calculated value and the current
// value from the array. On the first call, initial is used as the "previously
// calculated value."
//
//     nums := []int{2, 2, 12}
//     divide := func (previous int, current int) int { return previous / current }
//     result := ReduceRight(nums, divide, 144) // 3
func ReduceRight[V comparable](ary []V, reducer func(previous V, current V) V, initial V) V {
	result := initial

	for i := len(ary) - 1; i >= 0; i-- {
		result = reducer(result, ary[i])
	}

	return result
}

// Returns true if any element of the array satisfies the test function.
//
//     nums := []int{1,2,3}
//     isEven := func (v int) { return v % 2 == 0 }
//     hasEven := Any(nums, isEven) // true
func Any[V comparable](ary []V, test func(value V) bool) bool {
	for _, v := range ary {
		if test(v) {
			return true
		}
	}

	return false
}

// Returns true if every element of the array satisfies the test function.
//
//     nums := []int{2,4,6}
//     isEven := func (v int) { return v % 2 == 0 }
//     hasEven := Every(nums, isEven) // true
func Every[V comparable](ary []V, test func(value V) bool) bool {
	for _, v := range ary {
		if test(v) == false {
			return false
		}
	}

	return true
}

// Returns a new array containing only the elements in source that
// do not exist in other.
//
//     a := []int{1, 2, 3}
//     b := []int{2}
//     diff := Difference(a, b) // []int{1, 3}
func Difference[A comparable](source []A, other []A) []A {
	otherMap := make(map[A]bool)

	for _, v := range other {
		otherMap[v] = true
	}

	result := make([]A, 0)
	for _, v := range source {
		if _, found := otherMap[v]; !found {
			result = append(result, v)
		}
	}

	return result
}
