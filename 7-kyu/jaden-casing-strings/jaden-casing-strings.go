package kata
​
import "strings"
​
func ToJadenCase(str string) (res string) {
  for i, c := range str {
    if i == 0 || str[i-1] == ' ' {
      res += strings.ToUpper(string(c))
    }else {
      res += string(c)
    }
  }
  return
}