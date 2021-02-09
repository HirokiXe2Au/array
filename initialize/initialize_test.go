package initialize

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestLinearLinearNumber(t *testing.T) {
  actual := LinearNumber(10, -10)
  expected := []int{10 9 8 7 6 5 4 3 2 1 0 -1 -2 -3 -4 -5 -6 -7 -8 -9}

  assert.Equal(t, expected, actual)
}

func TestLinearLinearUpperCapital(t *testing.T) {
  actual := LinearUpperCapital(10, -10)
  expected := []string{"K", "J", "I", "H", "G", "F", "E", "D", "C", "B", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}

  assert.Equal(t, expected, actual)
}

func TestLinearLinearLowerCapital(t *testing.T) {
  actual := LinearLowerCapital(10, -10)
  expected := []string{"k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

  assert.Equal(t, expected, actual)
}