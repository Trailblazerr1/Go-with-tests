package arrays

// import "fmt"

func ArraySum(nos [5]int) int {
	sum := 0
	//range takes in array and return two values index and no.
	//we are ignoring index by using _ , so for is looping on around no
	for _, no := range nos {
		sum += no
	}
	return sum
}

func SliceSum(nos []int) (sum int) {
	for _, no := range nos {
		sum += no
	}
	return
}

// use ... to accept unlimite no of arguments
func SliceSumAll(nos ...[]int) []int {
	// lengthOfOuterSlice := len(nos)
	// //use make() to create new slice of particular len
	// ans := make([]int, lengthOfOuterSlice)
	// //range can't be nested
	// for i, outer := range nos {
	// 	ans[i] = SliceSum(outer)
	// }
	//------Above code can be rewritten as below
	//append returns the slice along with the second param appended
	var ans []int
	for _, outer := range nos {
		if len(outer) == 0 {
			ans = append(ans, 0)
		} else {
			//else can't be in next line
			ans = append(ans, SliceSum(outer))
		}
	}
	return ans
}

func SliceSumTails(nos ...[]int) []int {
	var ans []int
	for _, outer := range nos {
		if len(outer) == 0 {
			ans = append(ans, 0)
		} else {
			ans = append(ans, SliceSum(outer[1:]))
		}
	}
	return ans
}
