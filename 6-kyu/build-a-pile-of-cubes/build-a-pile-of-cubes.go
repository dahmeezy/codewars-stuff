package kata
​
func FindNb(m int) int {
  for n := 1 ; m > 0 ; n++ {
    m -= n*n*n
    if m == 0 { return n }
  }
  return -1
}