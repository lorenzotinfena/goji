package sort

func LongestIncreasingSubsequence[T any](v []T, prior func(a, b T) bool, stronglyIncreasing bool) int {
	dp := make([]T, 0)
	for i := 0; i < len(v); i++ {
		var j int
		if stronglyIncreasing {
			j = LowerBound(dp, v[i], prior)
		} else {
			j = UpperBound(dp, v[i], prior)
		}
		if j == len(dp) {
			dp = append(dp, v[i])
		} else {
			dp[j] = v[i]
		}
	}
	return len(dp)
}
