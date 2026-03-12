package kata
​
import ("strings"
        "strconv"
        )
​
func countSheep(num int) string {
  // Your code here!
  var res strings.Builder
  
  for i := 1; i <= num ; i++ {
    
    res.WriteString(strconv.Itoa(i))
    
    res.WriteString(" sheep...")
    
  }
  
  return res.String()
  
}