package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := range haystack {
		if haystack[i] == needle[0] && i+len(needle) <= len(haystack) {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
