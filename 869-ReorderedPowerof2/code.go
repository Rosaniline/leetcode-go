package leetcode

import (
	"math"
	"reflect"
)

func reorderedPowerOf2(N int) bool {
	var (
		digitsOfN  	= getNumberOfDigits(N)
		counterN 	= getCounterOfDigits(N)
	)

	for power := 1; ; power *= 2 {
		digitsOfPower := getNumberOfDigits(power)

		if digitsOfPower == digitsOfN {
			counterPower := getCounterOfDigits(power)

			if reflect.DeepEqual(counterPower, counterN) {
				return true
			}

		} else if digitsOfPower > digitsOfN {
			break
		}
	}

	return false
}


func getNumberOfDigits(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func getCounterOfDigits(n int) map[int]int {
	counter := make(map[int]int)

	for ; n > 0; n /= 10 {
		counter[n%10] ++
	}

	return counter
}