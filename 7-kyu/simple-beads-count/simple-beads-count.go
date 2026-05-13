package kata
​
func CountRedBeads(n int) int {
  if n<2{
    return 0 // your code here
  }
  return 2*(n-1)
}