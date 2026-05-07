package kata
​
func Tribonacci(signature [3]float64, n int) []float64 {
    // your code here
//   handled edge case for when n=0
  if n == 0{
    return []float64{}
  }
//   handled for when n<=0
  if n <= 3{
    return signature[:n]
  }
//   storage box
  var res []float64
  
//   store the orginal values in box
  res = append(res, signature[:]...)
  
//   set the loop to run n-3 times because it already has a slice oflength 3
  for i := 1; i <= n-3; i++{
    
//     the loops runs for as long as res length is not n
    if len(res) != n{
      
//       to get the position to start summing from(last 3 numbers basically)
      p := len(res) - 3
      var sum float64
      
      for j := p; j < len(res); j++{
//         store the sum in a variable and append to res
        sum += res[j]
      }
      res = append(res,sum)
    }
  }
  
 return res
}