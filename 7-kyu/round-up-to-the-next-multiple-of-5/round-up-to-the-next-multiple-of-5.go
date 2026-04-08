package kata
​
​
func RoundToNext5(n int) int {
  for n % 5 != 0 {
    n++
  }
  return n
}