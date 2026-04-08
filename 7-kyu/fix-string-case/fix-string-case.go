package kata
​
import "strings"
​
func solve(str string) string {
  // Your code here and happy coding!
  var up int
  var low int
  for _,c:= range str{
    if c>='a'&&c<='z'{
      low++
    } else{
      up++
    }
  }
  var a string
  if up>low{
    a=strings.ToUpper(str)
  } else if low> up{
    a=strings.ToLower(str)
  } else{
    a=strings.ToLower(str)
  }
  return a
  
}