func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false;
    }

    var sig, res [26]int
    for i := 0; i < len(s); i++ {
        sig[s[i]-'a']++
        res[t[i]-'a']++
    }
    return sig == res
}
