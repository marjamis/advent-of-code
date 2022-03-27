package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLocationValid(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		array := [][]int{
			{1, 2, 3, 4, 5, 6},
			{11, 12, 13, 14, 15, 16},
			{21, 22, 23, 24, 25, 26},
		}

		t.Run("Too Low", func(t *testing.T) {
			assert.False(t, IsLocationValid(array, -1, -1))
			assert.False(t, IsLocationValid(array, 1, -1))
			assert.False(t, IsLocationValid(array, -1, 1))
		})

		t.Run("Too High", func(t *testing.T) {
			assert.False(t, IsLocationValid(array, 1, 3))
			assert.False(t, IsLocationValid(array, 6, 1))
			assert.False(t, IsLocationValid(array, 6, 3))
		})

		t.Run("Just Right", func(t *testing.T) {
			assert.True(t, IsLocationValid(array, 0, 0))

			assert.True(t, IsLocationValid(array, 1, 0))
			assert.True(t, IsLocationValid(array, 2, 0))
			assert.True(t, IsLocationValid(array, 3, 0))
			assert.True(t, IsLocationValid(array, 4, 0))
			assert.True(t, IsLocationValid(array, 5, 0))

			assert.True(t, IsLocationValid(array, 0, 1))
			assert.True(t, IsLocationValid(array, 0, 2))
		})
	})
}

func TestCopy2dArray(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		original := [][]int{{1, 2, 3}, {4, 5, 6}}
		copy := Copy2dArray(original)
		assert.Equal(t, copy, original)
		assert.NotSame(t, copy, original)
	})

	t.Run("int32", func(t *testing.T) {
		original := [][]int32{{1, 2, 3}, {4, 5, 6}}
		copy := Copy2dArray(original)
		assert.Equal(t, copy, original)
		assert.NotSame(t, copy, original)
	})

	t.Run("int64", func(t *testing.T) {
		original := [][]int64{{1, 2, 3}, {4, 5, 6}}
		copy := Copy2dArray(original)
		assert.Equal(t, copy, original)
		assert.NotSame(t, copy, original)
	})

	t.Run("float32", func(t *testing.T) {
		original := [][]float32{{1.0, 2.1, 3.2}, {4.3, 5.4, 6.5}}
		copy := Copy2dArray(original)
		assert.Equal(t, copy, original)
		assert.NotSame(t, copy, original)
	})

	t.Run("float64", func(t *testing.T) {
		original := [][]float64{{1.0, 2.1, 3.2}, {4.3, 5.4, 6.5}}
		copy := Copy2dArray(original)
		assert.Equal(t, copy, original)
		assert.NotSame(t, copy, original)
	})

	t.Run("string", func(t *testing.T) {
		original := [][]string{{"1", "2", "3"}, {"4", "5", "6"}}
		copy := Copy2dArray(original)
		assert.Equal(t, copy, original)
		assert.NotSame(t, copy, original)
	})
}

func TestRemoveItemsAtIndexes(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5, 6}
		updated := RemoveItemsAtIndexes(original, []int{0, 2, 5})
		expected := []int{2, 4, 5}
		assert.Equal(t, expected, updated)
	})

	t.Run("int32", func(t *testing.T) {
		original := []int32{1, 2, 3, 4, 5, 6}
		updated := RemoveItemsAtIndexes(original, []int{0, 2, 5})
		expected := []int32{2, 4, 5}
		assert.Equal(t, expected, updated)
	})

	t.Run("int64", func(t *testing.T) {
		original := []int64{1, 2, 3, 4, 5, 6}
		updated := RemoveItemsAtIndexes(original, []int{0, 2, 5})
		expected := []int64{2, 4, 5}
		assert.Equal(t, expected, updated)
	})

	t.Run("float32", func(t *testing.T) {
		original := []float32{1.0, 2.1, 3.2, 4.3, 5.4, 6.5}
		updated := RemoveItemsAtIndexes(original, []int{0, 2, 5})
		expected := []float32{2.1, 4.3, 5.4}
		assert.Equal(t, expected, updated)
	})

	t.Run("float64", func(t *testing.T) {
		original := []float64{1.0, 2.1, 3.2, 4.3, 5.4, 6.5}
		updated := RemoveItemsAtIndexes(original, []int{0, 2, 5})
		expected := []float64{2.1, 4.3, 5.4}
		assert.Equal(t, expected, updated)
	})

	t.Run("string", func(t *testing.T) {
		original := []string{"1", "2", "3", "4", "5", "6"}
		updated := RemoveItemsAtIndexes(original, []int{0, 2, 5})
		expected := []string{"2", "4", "5"}
		assert.Equal(t, expected, updated)
	})
}
