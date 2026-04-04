type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder
    for _, v := range strs { 
        sb.WriteString(strconv.Itoa(len(v)))
        sb.WriteByte('#')
        sb.WriteString(v)
    }
    return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
    var result []string
    for encoded != ""{
        idx := strings.IndexByte(encoded, '#')
        count, _ := strconv.Atoi(encoded[:idx])
        result = append(result, encoded[idx+1:idx+1+count])
        encoded = encoded[idx+1+count:]
    }
    return result
}
