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
	sorted := make([]int, 100)

	for i := 0; i < 100; i++ {
		sorted[i] = 2 * i
	}

	result := search(sorted, 0)

	expected_response := 0

	assert.Equal(t, expected_response, result, "The left-most element could not be found")
}

func TestFindLastElement(t *testing.T) {
	sorted := make([]int, 100)

	for i := 0; i < 100; i++ {
		sorted[i] = 2 * i
	}

	result := search(sorted, 198)

	expected_response := 99

	assert.Equal(t, expected_response, result, "The right-most element could not be found")
}

func TestDoNotFindUnexistingElement(t *testing.T) {
	sorted := make([]int, 100)

	for i := 0; i < 100; i++ {
		sorted[i] = 2 * i
	}

	result := search(sorted, 200)

	expected_response := -1

	assert.Equal(t, expected_response, result, "The unexisting element was found")
}
