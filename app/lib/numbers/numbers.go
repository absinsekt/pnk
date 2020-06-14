package numbers

import (
	"math"
)

// Robin returns looped value, limited by min and max
func Robin(value, min, max int64) int64 {
	diff := (max + 1) - (min - 1)

	if value < min {
		result := min - value
		offset := int64(math.Mod(float64(result), float64(diff)))

		return max - offset + 1
	} else if value > max {
		result := value - max
		offset := int64(math.Mod(float64(result), float64(diff)))

		return min + offset - 1
	}

	return value
}
