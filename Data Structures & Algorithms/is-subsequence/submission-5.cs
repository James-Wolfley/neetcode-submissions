public class Solution {
    public bool IsSubsequence(string s, string t) {
        if (String.IsNullOrWhiteSpace(s))
            return true;
        if (String.IsNullOrWhiteSpace(t))
            return false;
        int idx = 0;
        for (int i = 0; i < t.Length; i++){
            if (s[idx] == t[i])
                idx++;
                if (idx == s.Length)
                    return true;
        }
        return false;
    }
}