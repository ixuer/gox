package slice

func StringSliceInclude(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func IntSliceInclude(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Int32SliceInclude(slice []int32, item int32) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Int64SliceInclude(slice []int64, item int64) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Float64SliceInclude(slice []float64, item float64) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Float32SliceInclude(slice []float32, item float32) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
