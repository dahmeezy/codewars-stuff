package kata
​
import (
  "sort"
  "strings"
)
​
func TwoSort(arr []string) string {
  var res strings.Builder
  sort.Strings(arr)
  fword := arr[0]
  for i, v := range arr[0] {
    res.WriteRune(v)
    if i != len(fword)-1 {
      res.WriteString("***")
    }
​
  }
​
  return res.String()
}
​