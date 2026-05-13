package kata
​
func GetMiddle(s string) string {
  //Code goes here!
  var r int
  if len(s)%2!=0{
    r=int(len(s)/2)
  } else{
    r=len(s)/2
  }
  var res string
  for i:=0;i<len(s);i++{
    if len(s)%2!=0{
      res=string(s[r])
    } else{
      res=string(s[r-1:r+1])
    }
  }
  return res
}