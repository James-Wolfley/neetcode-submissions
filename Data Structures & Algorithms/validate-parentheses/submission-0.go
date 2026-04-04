func isValid(s string) bool {
    stack := &Stack{}
    closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if open, exists := closeToOpen[c]; exists {
            if stack.Empty() {
                return false
            }
            top, _ := stack.Pop()
            if top != open {
                return false
            }
        } else {
            stack.Push(c)
        }
    }

    return stack.Empty()
}

type Stack []rune

func (s *Stack) Push(v rune) {
    *s = append(*s, v)
}

func (s *Stack) Pop() (rune, bool) {
    if s.Empty() {
        return ' ', false
    }
    l := len(*s)
    val := (*s)[l-1]
    *s = (*s)[:l-1] // actually shrink the slice
    return val, true
}

func (s *Stack) Empty() bool {
    return len(*s) == 0
}