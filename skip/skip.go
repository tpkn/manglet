package skip

import (
	"strconv"
	"strings"
)

// Rows converts columns string to a convenient map
func Rows(s string) map[int]string {
	result := map[int]string{}
	for _, v := range parseRows(s) {
		if _, found := result[v]; !found && v != 0 {
			result[v] = ""
		}
	}
	return result
}

func parseRows(rows string) []int {
	var result []int
	for _, v := range strings.Split(rows, ",") {
		r, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		result = append(result, r)
	}
	return result
}
