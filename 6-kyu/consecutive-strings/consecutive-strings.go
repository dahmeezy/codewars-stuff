package kata
​
import "strings"
​
    // your code
func LongestConsec(strarr []string, k int) string {
  // your code
  n := len(strarr)
  
  if n == 0 || k <= 0 || k > n {
    return ""
  }
  
  // a slice of string to store the concatenated words
  var concWords []string
  for i := 0; i <= len(strarr)-k; i++ {
    nWord := ""
    nWord += strings.Join(strarr[i:i+k], "")
    concWords = append(concWords, nWord)
  }
​
  longestString := ""
  longestStringLen := 0
​
  for j := range concWords {
    wordLen := len(concWords[j])
    // check if current wordlength is greater than our longest string and replace it with current word length if its true
    if wordLen > longestStringLen {
      longestStringLen = wordLen
      longestString = concWords[j]
    }
  }
  return longestString
}