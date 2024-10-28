package slice

// StringSliceUnique Remove duplicate data
func StringSliceUnique(slice []string) []string {
	var result []string
	tempMap := map[string]byte{}
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// IntSliceUnique Remove duplicate data
func IntSliceUnique(slice []int) []int {
	var result []int
	tempMap := map[int]byte{}
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// Int64SliceUnique Remove duplicate data
func Int64SliceUnique(slice []int64) []int64 {
	var result []int64
	tempMap := map[int64]byte{}
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// Int32SliceUnique Remove duplicate data
func Int32SliceUnique(slice []int32) []int32 {
	var result []int32
	tempMap := map[int32]byte{}
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// Float32SliceUnique Remove duplicate data
func Float32SliceUnique(slice []float32) []float32 {
	var result []float32
	tempMap := map[float32]byte{}
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

// Float64SliceUnique Remove duplicate data
func Float64SliceUnique(slice []float64) []float64 {
	var result []float64
	tempMap := map[float64]byte{}
	for _, e := range slice {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}
