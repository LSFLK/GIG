package commons

import "strings"

/**
match strings using Levenshtein distance
source: https://en.wikipedia.org/wiki/Levenshtein_distance
translated from C to Go
 */
func StringMatchPercentage(string1 string, string2 string) int {
	n := len(string1)
	m := len(string2)
	string1 = strings.ToLower(string1)
	string2 = strings.ToLower(string2)
	maxDifference := Maximum([]int{n, m})
	// create two work vectors of integer distances
	var (
		v0, v1 []int
	)
	// initialize v0 (the previous row of distances)
	// this row is A[0][i]: edit distance for an empty s
	// the distance is just the number of characters to delete from t

	for i := 0; i <= n; i++ {
		v0 = append(v0, i)
		v1 = append(v1, 0)
	}
	for j := 0; j < m; j++ {
		// calculate v1 (current row distances) from the previous row v0

		// first element of v1 is A[i+1][0]
		//   edit distance is delete (i+1) chars from s to match empty t
		v1[0] = j + 1

		// use formula to fill in the rest of the row
		for k := 0; k < n; k++ {
			// calculating costs for A[i+1][j+1]
			deletionCost := v0[k+1] + 1
			insertionCost := v1[k] + 1
			substitutionCost := 0
			if j < n && k < m && string1[j] == string2[k] {
				substitutionCost = v0[k]
			} else {
				substitutionCost = v0[k] + 1
			}
			min := Minimum([]int{deletionCost, insertionCost, substitutionCost})
			v1[k+1] = min
		}

		// copy v1 (current row) to v0 (previous row) for next iteration
		v0, v1 = v1, v0
		// after the last swap, the results of v1 are now in v0
	}

	return (maxDifference - v0[n]) * 100 / maxDifference
}

/**
Return a boolean value by matching two strings based on a given tolerance
 */
func StringsMatch(string1 string, string2 string, tolerance int) bool {
	matchPercent := StringMatchPercentage(string1, string2)
	return matchPercent >= tolerance
}
