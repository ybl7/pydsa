package main

func NestedListWeightedSum(arr []interface{}, sum, depth int) int {
	for _, e := range arr {
		if IsInt(e) {
			sum += e.(int) * depth
		} else {
			sum += NestedListWeightedSum(e.([]interface{}), 0, depth+1)
		}
	}
	return sum
}

// The list can contain both integers and slices of integers so the best we can do in go is []interface
func IsInt(e interface{}) bool {
	switch e.(type) {
	case int:
		return true
	}
	return false
}
