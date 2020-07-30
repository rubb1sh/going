package test

//
//func TestWB(t *testing.T) {
//
//}

// 739.

func dailyTemperatures(T []int) []int {

	type node struct {
		Day   int
		Value int
	}

	stack := make([]node, len(T))
	stack[0] = node{0, T[0]}
	top := 0

	res := make([]int, len(T))

	for i := 1; i < len(T); i++ {
		if stack[top].Value < T[i] {
			for top >= 0 && stack[top].Value < T[i] {
				res[stack[top].Day] = i - stack[top].Day
				top--
			}
		}
		top++
		stack[top] = node{i, T[i]}
	}

	for top >= 0 {
		res[stack[top].Day] = 0
		top--
	}

	return res

}
