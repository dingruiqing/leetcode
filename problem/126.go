package problem

func LadderLength2(beginWord string, endWord string, wordList []string) int {
	n := len(wordList)
	m := len(beginWord)
	if n == 0 || m != len(endWord) {
		return 0
	}
	wordDict := make(map[string]bool)
	for i := 0; i < n; i++ {
		wordDict[wordList[i]] = true
	}
	if _, ok := wordDict[endWord]; !ok {
		return 0
	}
	if _, ok := wordDict[beginWord]; ok {
		delete(wordDict, beginWord)
	}
	queue := []string{beginWord}
	sum := 0
	for len(queue) != 0 {
		tmp := make([]string, 0)
		for _, v := range queue {
			for i := 0; i < m; i++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := ""
					switch i {
					case 0:
						newWord = string(j) + v[1:]
					case m - 1:
						newWord = v[:m-1] + string(j)
					default:
						newWord = v[:i] + string(j) + v[i+1:]
					}
					if newWord == endWord {
						return sum + 2
					}
					if _, ok := wordDict[newWord]; ok {
						tmp = append(tmp, newWord)
						delete(wordDict, newWord)
					}
				}
			}
		}
		sum++
		queue = tmp
	}
	return 0
}
