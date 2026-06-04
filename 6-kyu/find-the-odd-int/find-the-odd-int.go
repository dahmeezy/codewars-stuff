package kata
​
func FindOdd(seq []int) int {
  if len(seq) == 1 {
    return seq[0] // your code here
  }
​
  var res int
​
  mp := make(map[int]int)
​
//   populate the map with the values in the sequence
  for _, n := range seq {
    mp[n]++
  }
​
//   get the key of any odd value
  for key, value := range mp {
    if value%2 != 0 {
      res = key
    }
  }
​
  return res
}