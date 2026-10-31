package arraypkg

// ArraySlice returns a subset of the given string slice.
// It keeps the slice safe by checking the start and end values.
func ArraySlice(items []string, start, end int) []string {
	// Prevent negative indexes.
	if start < 0 {
		start = 0
	}
	// Do not go beyond the slice length.
	if end > len(items) {
		end = len(items)
	}
	// Prevent invalid ranges.
	if start > end {
		return []string{}
	}
	return items[start:end]
}
