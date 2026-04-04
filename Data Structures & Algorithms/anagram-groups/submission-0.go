func groupAnagrams(strs []string) [][]string {
    lists := make(map[string][]string)
    for _, v := range strs {
        b := []byte(v)
        sort.Slice(b, func(i,j int) bool {return b[i] < b[j]})
        temp_string := string(b)
        lists[temp_string] = append(lists[temp_string], v)
    }

    result := make([][]string, 0)
    for _, v := range lists{
        result = append(result, v)
    }
    return result
}
