package kata
​
func SumCubes(n int) int {
  // your code here
  res:=0
  for i:=n;i>=0;i--{
   c:= i*i*i
    res+=c
  }
  return res
}