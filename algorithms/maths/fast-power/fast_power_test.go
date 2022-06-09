package fast_power

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastPower(t *testing.T) {
	a := uint32(2)

	for i := 0; i < 100000; i++ {
		n, ok := fast_power(a, i)

		if ok != nil || n != uint32(math.Pow(float64(a), float64(i))) {
			t.Error()
		}
	}
}

func TestFastPowerEquiClass0(t *testing.T) {
	a := uint32(2)
	exp := 0
	expected_response := uint32(1)

	result, _ := fast_power(a, exp)

	assert.Equal(t, expected_response, result, "Every number to the power of 0 must be 1")
}

func TestSlowPowerEquiClass0(t *testing.T) {
	a := uint32(2)
	exp := 0
	expected_response := uint32(1)

	result, _ := slow_power(a, exp)

	assert.Equal(t, expected_response, result, "Every number to the power of 0 must be 1")
}

func TestSlowPower(t *testing.T) {
	a := uint32(2)

	for i := 0; i < 100000; i++ {
		n, ok := slow_power(a, i)

		if ok != nil || n != uint32(math.Pow(float64(a), float64(i))) {
			t.Error()
		}
	}
}

func BenchmarkFastPower(b *testing.B) {
	a := uint32(2)

	for i := 0; i < b.N; i++ {
		fast_power(a, 100000)
	}
}

func BenchmarkSlowPower(b *testing.B) {
	a := uint32(2)

	for i := 0; i < b.N; i++ {
		slow_power(a, 100000)
	}
}
