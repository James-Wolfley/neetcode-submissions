func isPalindrome(s string) bool {
    str := strings.ReplaceAll(s, " ", "")
    str = strings.ReplaceAll(str, "?", "")
    str = strings.ReplaceAll(str, ".", "")
    str = strings.ReplaceAll(str, ",", "")
    str = strings.ReplaceAll(str, "!", "")
    str = strings.ReplaceAll(str, "'", "")
    str = strings.ReplaceAll(str, ":", "")
    str = strings.ReplaceAll(str, ";", "")
    str = strings.ToLower(str)
    fmt.Println(str)
    for i := 0; i < len(str); i++{
        if str[i] != str[len(str)-i-1]{
            return false
        }
    }
    return true
}
