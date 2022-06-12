package rselect

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSelect(t *testing.T) {
	arr := []int{5, 2, 6, 8, 3, 1}

	i0 := RSelect(arr, 6, 0)
	i1 := RSelect(arr, 6, 1)
	i2 := RSelect(arr, 6, 2)
	i3 := RSelect(arr, 6, 3)
	i4 := RSelect(arr, 6, 4)
	i5 := RSelect(arr, 6, 5)
	if i0 != 1 &&
		i1 != 2 &&
		i2 != 3 &&
		i3 != 5 &&
		i4 != 6 &&
		i5 != 8 {
		fmt.Println("Error: ", i0, i1, i2, i3, i4, i5)
		t.Error()
	}
}

func TestLimitValue0(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	len := 5
	k_element := 0
	expected_response := 1

	result := RSelect(arr, len, k_element)

	assert.Equal(t, expected_response, result, "Element of order 0 should be the smallest element")
}

func TestLimitValueMinus1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	len := 5
	k_element := -1
	expected_response := 1

	result := RSelect(arr, len, k_element)

	assert.Equal(t, expected_response, result, "Element of order -1 should be the smallest element")
}

func TestLimitValue1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	len := 5
	k_element := 1
	expected_response := 2

	result := RSelect(arr, len, k_element)

	assert.Equal(t, expected_response, result, "Element of order 1 should be the second smallest element")
}

func TestLimitValueNMinus1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	len := 5
	k_element := 4
	expected_response := 5

	result := RSelect(arr, len, k_element)

	assert.Equal(t, expected_response, result, "Element of order N - 1 should be the largest element")
}

func TestLimitValueN(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	len := 5
	k_element := 5
	expected_response := 5

	result := RSelect(arr, len, k_element)

	assert.Equal(t, expected_response, result, "Element of order N should be the largest element")
}
