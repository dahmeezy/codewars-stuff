package kata
​
func MakeNegative(x int) int {
  if x>0{
    return -x
  } else if x< 0 {
    return x
  }else{
     return 0
  }
}