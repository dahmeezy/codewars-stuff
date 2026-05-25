package kata
​
import (
  "strconv"
)
func PrinterError(s string) string {
  l:=strconv.Itoa(len(s))
  res:=0  
  
  for _,c:=range s{
    if c>'a'&&c>'m'{
      res++
    }
  }
     n:=strconv.Itoa(res)
  sc:=n +"/"+l
      return sc
​
}
​