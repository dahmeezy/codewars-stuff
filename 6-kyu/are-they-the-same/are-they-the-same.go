package kata
​
import "sort"
func Comp(array1 []int, array2 []int) bool {
    // your code
  var same bool
  
//  catch when any of the array is empty
  if array1==nil||array2==nil{
    return false
  }  
//   makes sure that the length of both arrays are equal
  if len(array1) != len(array2) {
    return false
}
//   squares everything in the first array
  for i,n:=range array1{
    array1[i]=n*n
  }
//   sort the arrays
  sort.Ints(array1)
  sort.Ints(array2)
  
  var array3 []int
//   store the common ones
 for j, v := range array2 {
    if array1[j] == array2[j] {
      array3 = append(array3, v)
    }
  }
// check if its same length as any of the 2 previous array and choose your bool
    if len(array3)==len(array1){
      same=true
    }else {
      same=false
    }
  return same
  
}