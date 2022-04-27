package functional

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Simple add10 to int array", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := []int{11, 12, 13}
		actual := Map(nums, func(i int) int { return i + 10 })

		assert.Equal(t, expected, actual)
	})

	t.Run("Convert int array to string array", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		actual := Map(nums, func(i int) string { return strconv.Itoa(i) })

		assert.Equal(t, expected, actual)
	})
}

func TestFilter(t *testing.T) {
	t.Run("Filter 2 from int array", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := []int{1, 3}
		actual := Filter(nums, func(v int) bool { return v != 2 })

		assert.Equal(t, expected, actual)
	})
}

func TestZip(t *testing.T) {
	t.Run("Two string arrays", func(t *testing.T) {
		keys := []string{"name", "country"}
		values := []string{"John", "USA"}
		expected := map[string]string{"name": "John", "country": "USA"}

		assert.Equal(t, expected, Zip(keys, values))
	})

	t.Run("Integer as key, string as value", func(t *testing.T) {
		keys := []int{1, 2, 3}
		values := []string{"John", "Jack", "Jane"}
		expected := map[int]string{1: "John", 2: "Jack", 3: "Jane"}

		assert.Equal(t, expected, Zip(keys, values))
	})

	t.Run("String as key, any as value", func(t *testing.T) {
		keys := []string{"name", "age"}
		values := []any{"John", 22}
		expected := map[string]any{"name": "John", "age": 22}

		assert.Equal(t, expected, Zip(keys, values))
	})
}

func TestUnzip(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		var foundName, foundAge bool
		m := map[string]any{"name": "John", "age": 28}

		keys, values := Unzip(m)

		assert.Equal(t, len(m), len(keys), "Keys is not the proper length")
		assert.Equal(t, len(m), len(values), "Values is not the proper length")

		for i, k := range keys {
			switch k {
			case "name":
				foundName = true
				assert.Equal(t, "John", values[i])
			case "age":
				foundAge = true
				assert.Equal(t, 28, values[i])
			default:
				assert.Fail(t, fmt.Sprintf("unknown key: %s", k))
			}
		}

		assert.True(t, foundName, "Didn't find name in the returned keys!")
		assert.True(t, foundAge, "Didn't find age in the returned keys!")
	})
}

func TestReduce(t *testing.T) {
	t.Run("Sum values", func(t *testing.T) {
		f := func(previous int, current int) int { return previous + current }
		nums := []int{1, 2, 3}
		expected := 6

		assert.Equal(t, expected, Reduce(nums, f, 0))
	})
}

func TestReduceRight(t *testing.T) {
	t.Run("Divide values", func(t *testing.T) {
		nums := []int{2, 2, 12}
		divide := func(previous int, current int) int { return previous / current }
		result := ReduceRight(nums, divide, 144) // 3

		assert.Equal(t, 3, result)
	})
}

func TestAny(t *testing.T) {
	t.Run("Any true with a true value", func(t *testing.T) {
		ary := []bool{false, true, false}
		isTrue := func(v bool) bool { return v == true }

		assert.Equal(t, true, Any(ary, isTrue))
	})

	t.Run("Any true with no true values", func(t *testing.T) {
		ary := []bool{false, false, false}
		isTrue := func(v bool) bool { return v == true }

		assert.Equal(t, false, Any(ary, isTrue))
	})
}

func TestEvery(t *testing.T) {
	t.Run("All items are true", func(t *testing.T) {
		ary := []bool{true, true, true}
		isTrue := func(v bool) bool { return v == true }

		assert.Equal(t, true, Every(ary, isTrue))
	})

	t.Run("One item is false", func(t *testing.T) {
		ary := []bool{true, false, true}
		isTrue := func(v bool) bool { return v == true }

		assert.Equal(t, false, Every(ary, isTrue))
	})
}

func TestDifference(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		source := []int{1, 2, 3}
		other := []int{2}
		expected := []int{1, 3}

		assert.Equal(t, expected, Difference(source, other))
	})
}

func TestUnique(t *testing.T) {
	t.Run("Containing duplicate values", func(t *testing.T) {
		ary := []int{1, 1, 2, 3}
		expected := []int{1, 2, 3}
		actual := Unique(ary)

		sort.Ints(actual)

		assert.Equal(t, expected, actual)
	})
}
