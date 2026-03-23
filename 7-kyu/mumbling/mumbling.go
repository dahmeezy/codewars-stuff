package kata
​
import (
"strings"
)
​
func Accum(s string) string {
var res strings.Builder
  for i,c:=range s{
    for j:=0;j<=i;j++{
      if j==0 {
        res.WriteString(strings.ToUpper(string(c)))
      } else{
        res.WriteString(strings.ToLower(string(c)))
      }
    }
    if i==len(s)-1{
      continue
    }else{
      res.WriteString("-")
    }
    
//     res.WriteString(strings.ToUpper(string(rune(int(c)-32))))
  }
  return res.String()
}