package kata
​
import(
  "strings"
)
​
func High(s string) string {
  // your code here
  var res []int
​
//   seperate words by space and put in a slice
  words := strings.Fields(s)
  for _, word := range words {
    var val int
    for _, c := range word {
//       gets the position in alphabet
      p := int(c) - 96
      val += p
    }
//     stores list of values according to each word
    res = append(res, val)
​
  }
  maxval := 0
  maxindex := 0
​
//   compare values and get index of the highest
  for i, v := range res {
    if v > maxval {
      maxval = v
      maxindex = i
    }
  }
// return the actual word
  return words[maxindex]
}
​