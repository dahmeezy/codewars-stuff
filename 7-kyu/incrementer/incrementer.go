package kata
​
​
func Incrementer(n []int) []int {
      // your code here
  
  if n==nil{
    return []int{}
  }
  for i,v:=range n{
    n[i]=v+(i+1)
  }
  for i,v:=range n{
    if v>9{
      n[i]=v%10
    }
  }
  return n
}
​