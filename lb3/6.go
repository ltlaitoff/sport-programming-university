package lb3

// Є два відсортованих за не зростанням масиви A[1,N] і B[1, M]. Отримати
// відсортований за не зростанням масив C[1, N+M], що складається з елементів масивів A
// і B ("злити" разом масиви A і B).

func Lab3Task6(a []int, b []int) []int {
	c := make([]int, len(a)+len(b))

	recursiveArrayMerge(a, b, c, 0, 0, 0)

	return c
}

func recursiveArrayMerge(a []int, b []int, c []int, i int, j int, x int) {
	if i >= len(a) && j >= len(b) {
		return
	}

	if i >= len(a) {
		c[x] = b[j]
		recursiveArrayMerge(a, b, c, i, j+1, x+1)
		return
	}

	if j >= len(b) {
		c[x] = a[i]
		recursiveArrayMerge(a, b, c, i+1, j, x+1)
		return
	}

	if a[i] > b[j] {
		c[x] = a[i]
		recursiveArrayMerge(a, b, c, i+1, j, x+1)
		return
	}

	c[x] = b[j]
	recursiveArrayMerge(a, b, c, i, j+1, x+1)
}
