package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestItertoolsProduct(t *testing.T) {
	first := []string{"a"}
	second := []string{"d", "e"}
	result := ItertoolsProduct(first, second)
	assert.ElementsMatch(t, *result, [][2]string{[2]string{"a", "d"}, [2]string{"a", "e"}} )

}
