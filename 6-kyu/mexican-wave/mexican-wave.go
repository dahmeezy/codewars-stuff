package kata
​
import "unicode"
​
func wave(words string) []string {
  // Your code here and happy coding!
​
  var res []string
​
// handles cases when empty string is passed
  if words == "" {
    return []string{}
  }
// convert word into a slice of runes
  runes := []rune(words)
  
  for i, c := range runes {
//     create a copy of rune that resets after each iteration
    var dup []rune
    dup = append(dup, runes...)
//     handles string with space between
    if c == ' ' {
      continue
    }
​
//     uppercase the current character and append to result
    dup[i] = unicode.ToUpper(dup[i])
​
    res = append(res, string(dup))
  }
  return res
}
  