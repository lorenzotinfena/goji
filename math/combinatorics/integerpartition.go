package combinatorics

// Computes integer partitions, each partition is weakly increasing
func IntegerPartitions(n int) [][]int {
	stack := []int{}
	partitions := [][]int{}
	var computePartitions func(remaining int, maxLength int)
	computePartitions = func(remaining int, maxLength int) {
		if remaining == 0 {
			partitions = append(partitions, append([]int{}, stack...))
			return
		}
		for i := 1; i <= maxLength && i <= remaining; i++ {
			stack = append(stack, i)
			computePartitions(remaining-i, i)
			stack = stack[:len(stack)-1]
		}
	}
	computePartitions(n, n)
	return partitions
}
