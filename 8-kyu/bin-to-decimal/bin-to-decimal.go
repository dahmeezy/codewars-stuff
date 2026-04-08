package kata
 
import "strconv"
​
func BinToDec(bin string) int {
      // your code 
  dec,_:=strconv.ParseInt(bin,2,32)
  return int(dec)
}
​