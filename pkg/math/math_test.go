package mathutils

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNextNumber_Valid(t *testing.T) {
	tests := []struct {
		name     string
		nums     []uint64
		expected uint64
	}{
		{
			name:     "ReturnsNextNumber_WhenNumsIsValid",
			nums:     []uint64{1, 2, 3},
			expected: 4,
		},
		{
			name:     "Returns1_WhenNumsIsEmpty",
			nums:     []uint64{},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			next, err := GenerateNextNumber(tt.nums)

			assert.Equal(t, tt.expected, next)
			assert.Nil(t, err)
		})
	}
}

func TestGenerateNextNumber_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		nums     []uint64
		expected error
	}{
		{
			name:     "ReturnsError_WhenNumsIsNil",
			nums:     nil,
			expected: ErrNumsMustNotBeNil,
		},
		{
			name:     "ReturnsError_WhenNumsContainsMaxValue",
			nums:     []uint64{math.MaxUint64},
			expected: ErrMaxValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			next, err := GenerateNextNumber(tt.nums)

			assert.Equal(t, uint64(0), next)
			assert.Error(t, err)
			assert.ErrorIs(t, tt.expected, err)
		})
	}
}
