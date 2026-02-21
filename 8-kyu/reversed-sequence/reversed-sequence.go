package kata
​
func ReverseSeq(n int) []int {
  var ns []int
  for i:= n ; i > 0; i--{
    ns = append(ns, i)
  }
  return ns
}