package kata
​
​
func Xbonacci(signature []int, n int) []int {
//   solved like tribonacci, tailored for n instead of 3
  h:=len(signature)
  
  if h==0{
    return []int{}
  }
  if n==0{
    return []int{}
  }
  if h>n{
    return signature[:n]
  }
  
  var res []int
  
  res=append(res,signature...)
  
  for i:=1;i<=n-h;i++{
    if len(res) != n{
      
      p:=len(res)-h
      var sum int
      for j:=p;j<len(res);j++{
        
        sum+=res[j]
      }
      res=append(res,sum)
    }
  }
  return res
  
}
​