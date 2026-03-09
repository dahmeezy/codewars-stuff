package kata
​
​
import (
  "strconv"
)
​
func Digitize(n int) []int {
  // your code here
  a := strconv.Itoa(n)
  var result []int
  for i := len(a) - 1; i >= 0; i--{
    done, _ := strconv.Atoi(string(a[i]))
    result = append(result, done)
​
  }
​
  return result
}
​