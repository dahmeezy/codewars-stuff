package kata
​
​
func RepeatStr(rep int, val string) string {
 if rep > 0 {
   var new_str string
   for i := 1; i <= rep; i++ {
     new_str += val
   }
   return new_str
 }
  return ""
 }