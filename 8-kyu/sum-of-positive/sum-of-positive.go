package kata
​
func PositiveSum(numbers []int) int {
 var res int
  for _,v:=range numbers{
    if v>0{
      res+=v
    }
  }
  return res
}
​