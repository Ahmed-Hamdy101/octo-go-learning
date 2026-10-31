package arraypkg

func ArraySlice(items []string, start, end int) []string {
	if start < 0 {
		start = 0
	}
	if end > len(items) {
		end = len(items)
	}
	if start > end {
		return []string{}
	}
	return items[start:end]
}
