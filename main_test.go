package functional

import (
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
