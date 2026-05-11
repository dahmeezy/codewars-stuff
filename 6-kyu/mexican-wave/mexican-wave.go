package kata
​
import "strings"
​
func wave(s string) []string {
  // Your code here and happy coding!
  res := []string{}
  if len(s) == 0{
    return []string{}
  }
  for i, v := range s {
    if v == ' ' {
      continue
    }
    if i > 0 {
      word := s[0:i] + strings.ToUpper(s[i:i+1]) + s[i+1:]
      res = append(res, word)
    } else {
      word := strings.ToUpper(s[i:i+1]) + s[i+1:]
      res = append(res, word)
    }
  }
  return res
}