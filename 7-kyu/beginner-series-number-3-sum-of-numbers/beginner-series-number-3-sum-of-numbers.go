package kata
​
​
func GetSum(a, b int) int {
  var res int
  
  if a>b{
    for i:=b;i<=a;i++{
      res+=i
    }
    return res
  }
  
  if b>a{
    for i:=a;i<=b;i++{
      res+=i
    }
    return res
  }
  return a
}