package kata
​
import "strings"
​
func solve(slice []string) []int {
  // Your code here and happy coding!
​
  var res []int
​
  for ind, word := range slice {
    slice[ind] = strings.ToLower(word)
​
    count := 0
    for i, c := range slice[ind] {
      if int(c - 'a' + 1) == i+1 {
        count++
      }
    }
    res = append(res, count)
​
  }
  return res
}