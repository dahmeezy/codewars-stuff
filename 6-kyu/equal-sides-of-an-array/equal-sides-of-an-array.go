package kata
​
​
func FindEvenIndex(arr []int) int {
  
// checks if slice contains nothging
  if len(arr) == 0 {
    return -1
  }
  
  for i := range arr {
//     creates variables that resets every single loop
    sumLeft := 0
    sumRight := 0
//     sum elements on either side of current element
    for _, nl := range arr[:i] {
      sumLeft += nl
    }
    for _, nr := range arr[i+1:] {
      sumRight += nr
    }
//     return the index of where elements of either sides sre equal
    if sumLeft == sumRight {
      return i
    } 
    
  }
  return -1
}
​