package kata
​
func SquareSum(numbers []int) int {
  var res int
    // loop through the array of numbers
  for _,n:=range numbers{
//     square each of them 
    s:=n*n
//     store
    res+=s
  }
//   return sum
  return res
  
}