package bs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	sorted := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		sorted[i] = 2 * i
	}

	for i := 0; i < 10000; i++ {
		index := search(sorted, 2*i)

		if index != i {
			fmt.Println(index)
			t.Error()
		}
	}

	if search(sorted, 3) != -1 {
		t.Error()
	}
}

func TestFindFirstElement(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	result := search(sorted, 1)
	expected_response := 0

	assert.Equal(t, expected_response, result, "The left-most element could not be found")
}

func TestFindLastElement(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	result := search(sorted, 5)
	expected_response := 4

	assert.Equal(t, expected_response, result, "The right-most element could not be found")
}

func TestFindDeepestElement(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	result := search(sorted, 2)
	expected_response := 1

	assert.Equal(t, expected_response, result, "The deepest element could not be found")
}

func TestFindEasiestElement(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	result := search(sorted, 3)
	expected_response := 2

	assert.Equal(t, expected_response, result, "The easiest element could not be found")
}

func TestDoNotFindUnexistingElement(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	result := search(sorted, 6)
	expected_response := -1

	assert.Equal(t, expected_response, result, "The unexisting element was found")
}
