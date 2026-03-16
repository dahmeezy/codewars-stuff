package kata
​
​
​
func MinMax(arr []int) [2]int {
  res := [2]int{}
​
  min := arr[0]
  max := arr[0]
  for i := 1; i < len(arr); i++ {
    if arr[i] >= max {
      max = arr[i]
    }
​
    if arr[i] <= min {
​
      min = arr[i]
    }
    res[0] = min
    res[1] = max
​
  }
  return res
}
​
​