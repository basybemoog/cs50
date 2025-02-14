package week_6

func longestCommonPrefix(strs []string) string {
	var output string
	if len(strs) < 2 {
		return strs[0]
	}
	var length = len(strs[0])
	for i := 0; i < len(strs)-1; i++ {
		if length > len(strs[i+1]) {
			length = len(strs[i+1])
		}
	}
	check := false
	if len(strs[0]) > 0 {
		for i := 0; i < length; i++ {

			for j := 0; j < len(strs)-1; j++ {
				if strs[j][i] == strs[j+1][i] {
					check = true
				} else {
					check = false
					break
				}
			}
			if check {
				output += string(strs[0][i])
			} else {
				break
			}
		}
	}

	return output
}
