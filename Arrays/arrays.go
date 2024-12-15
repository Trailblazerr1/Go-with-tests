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
	lengthOfOuterSlice := len(nos)
	//use make() to create new slice of particular len
	ans := make([]int, lengthOfOuterSlice)
	//range can't be nested
	for i, outer := range nos {
		ans[i] = SliceSum(outer)
	}
	return ans
}
