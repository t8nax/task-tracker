package mathutils

import (
	"errors"
	"math"
)

var ErrMaxValue = errors.New("nums contains max value of uint64")
var ErrNumsMustNotBeNil = errors.New("nums must not be nil")

func GenerateNextNumber(nums []uint64) (uint64, error) {
	if nums == nil {
		return 0, ErrNumsMustNotBeNil
	}

	var next uint64 = 0

	for _, num := range nums {
		if num == math.MaxUint64 {
			return 0, ErrMaxValue
		}

		if num > next {
			next = num
		}
	}

	return next + 1, nil
}
