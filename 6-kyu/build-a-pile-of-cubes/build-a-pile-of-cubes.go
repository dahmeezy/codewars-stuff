package kata
​
import (
"math")
​
func FindNb(m int) int {
  // your code
//   n=(-1 + \sqrt{1+(8*\sqrt{m})})/2 so thats the calculation i'm doing here
  s := 8 * math.Sqrt(float64(m))
  n := int((math.Sqrt(s+1) - 1) / 2)
  var sum int
  var count int
//   reverse loop 
  for i := n; i > 0; i-- {
//     cube and store
    sum+=int(math.Pow(float64(i), 3))
    count++
}
//   return number of cubes only if the sum equals m
  if sum == m {
    return count
//     if there is no such n we should return -1(if the sum is not equal to m)
  } else {
    return -1
  }
​
  }
  
​