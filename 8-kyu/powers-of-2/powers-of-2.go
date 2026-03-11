package kata
​
func PowersOfTwo(n int) []uint64 {
  
  var res []uint64
  
  // your code goes here
  for i:=0;i<=n;i++{
    
//     1<<i = 1*2^i(bit-wise shift operator)
    v:=1<<i
    
    res=append(res, uint64(v))
    
  }
  return res
  
}