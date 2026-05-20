package kata
​
func Seven(n int64) []int {
    // your code
  d:=int(n)
  var res []int
  var count int
  for d>99{
    lastDigit:=d%10
    remDigits:=d/10
    newDigit:=remDigits-2*(lastDigit)
    count++
    d=newDigit
    }
   res=append(res,d)
   res=append(res,count)
  return res
    
}
  