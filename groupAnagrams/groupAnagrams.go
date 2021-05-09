package main

func main() {
}

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	strMap := make(map[string][]string, 0)
	for _, v := range strs {
		sortStr := sortStringByAlpha(v)
		strMap[sortStr] = append(strMap[sortStr], v)
	}
	for _, v := range strMap {
		ret = append(ret, v)
	}
	return ret

}
func sortStringByAlpha(s string) string {
	b := []byte(s)
	n := len(s)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if b[i] > b[j] {
				tmp := b[i]
				b[i] = b[j]
				b[j] = tmp
			}
		}
	}
	ret := string(b)
	return ret
}
