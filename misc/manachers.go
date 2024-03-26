package misc

// To get longest odd palidrome starting centered in i: v[2*i+1]
// To get longest even palidrome starting centered between i and i+1: v[2*i+2]
func ManachersAlgorithm[T comparable](v []T) []int {
	supp := make([]T, len(v)*2+1)
	var foo T
	for i := 0; i < len(v); i++ {
		supp[2*i] = foo
		supp[2*i+1] = v[i]
	}
	res := make([]int, len(v)*2+1)
	l := 1
	r := 1
	for i := 1; i < len(supp); i++ {
		if i > r {
			l = i
			r = i
		}
		if res[l+r-i] < r-i {
			res[i] = res[l+r-i]
		} else {
			res[i] = r - i
			for i+res[i]+1 < len(supp) && i-res[i]-1 >= 0 && supp[i+res[i]+1] == supp[i-res[i]-1] {
				res[i]++
			}
			l = i - res[i]
			r = i + res[i]
		}
	}
	return res
}
