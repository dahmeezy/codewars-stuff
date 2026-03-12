package kata
​
import "strings"
​
func GetCount(str string) (count int) {
  
  // Enter solution here
  for _,char := range str {
    
    if strings.ContainsAny(string(char),"aeiou"){
      
      count++
      
    }
  }
  return count
}