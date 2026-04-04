func minWindow(s string, t string) string {
    //move r till we have all of t then move l till we lose one and almost another, then repeat till r is = len
	//use freq to match it
	if t == "" {
		return ""
	}

	sig := make(map[rune]int)
	for _, c := range t {
		sig[c]++
	}

	have, need := 0, len(sig)
	res := []int{-1,-1}
	resLen := math.MaxInt32
	l := 0
	window := make(map[rune]int)
	for r := 0; r < len(s); r++ {
		c := rune(s[r])
		window[c]++

		if sig[c] > 0 && window[c] == sig[c] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				res = []int{l,r}
				resLen = r - l + 1
			}

			window[rune(s[l])]--
			if sig[rune(s[l])] > 0 && window[rune(s[l])] < sig[rune(s[l])] {
				have--
			}
			l++
		}
	}

	if res[0] == -1 {
		return ""
	}
	return s[res[0]:res[1]+1]

}
