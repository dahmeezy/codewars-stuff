package kata
​
func FindNb(m int) int {
  return FindNbRecursive(m, 1)
}
​
func FindNbRecursive(m int,i int) int {
  m -= (i*i*i)
  if m == 0 {
    return i
  }
  if m < 0 {
    return -1
  }
  return FindNbRecursive(m, i+1)
}