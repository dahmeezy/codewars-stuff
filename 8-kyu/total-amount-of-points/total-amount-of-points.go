package kata
import (
        "strconv")
​
func Points(games []string) int {
  var res int
  // your code here!
  for _,v:=range games{
    x,_:=strconv.Atoi(string(v[0]))
    y,_:=strconv.Atoi(string(v[2]))
    if x>y{
      res+=3
    } else if x<y{
      res+=0
    } else {
      res+=1
    }
  }
  return res
}