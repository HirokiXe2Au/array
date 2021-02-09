package search

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
  arr := [...] string{"Golang", "Java"}
  actual := linearSearch(3, [])
  expected := 8

  assert.Equal(t, expected, actual)
}`