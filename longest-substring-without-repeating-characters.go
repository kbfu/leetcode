package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	count := 0
	max := 0
	if len(s) > 1 {
		for i := range s {
			char := string(s[i])
			_, ok := m[char]
			if ok {
				if count > max {
					max = count
					count = 1
				}
			} else {
				count++
				m[char] = i
				if i == len(s)-1 && count > max {
					max = count
				}
			}
		}
	} else if len(s) == 1 {
		max = 1
	}

	return max
}

func main() {
	println(lengthOfLongestSubstring(" "))
	println(lengthOfLongestSubstring("abcabc"))
	println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring(""))
	println(lengthOfLongestSubstring("au"))
}
