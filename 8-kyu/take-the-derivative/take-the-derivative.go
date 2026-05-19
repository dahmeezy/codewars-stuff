package kata
​
import "strconv"
​
​
func Derive(coefficient, exponent int) string {
  
  p:=coefficient * exponent
  s:=strconv.Itoa(p)
  dif:=strconv.Itoa(exponent-1)
  return s+"x^"+dif
}
​