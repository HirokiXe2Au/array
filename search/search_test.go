package search

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestLinearSearchInt(t *testing.T) {
  arr := []int{10 9 8 7 6 5 4 3 2 1 0 -1 -2 -3 -4 -5 -6 -7 -8 -9}
  actual := LinearSearchInt(3, arr, 20)
  expected := 7

  assert.Equal(t, expected, actual)
}

func TestLinearSearchString(t *testing.T) {
  arr := []string{"K", "J", "I", "H", "G", "F", "E", "D", "C", "B", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
  actual := LinearSearchString("J", arr, 20)
  expected := 1

  assert.Equal(t, expected, actual)
}`