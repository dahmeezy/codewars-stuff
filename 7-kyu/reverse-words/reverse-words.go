package kata
​
import (
  "strings"
)
​
func ReverseWords(str string) string {
  var res strings.Builder
​
  words := strings.Split(str, " ")
​
  for i, word := range words {
    for i := len(word) - 1; i >= 0; i-- {
      res.WriteString(string(word[i]))
​
    }
    if i == len(words)-1 {
      continue
    }else{
      res.WriteString(" ")
    }
​
  }
  return res.String() // reverse those words
}