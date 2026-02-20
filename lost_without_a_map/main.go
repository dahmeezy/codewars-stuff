package kata

func Maps(x []int) []int {
  //create a new slice to store the new values
  var new_a []int
  //loop through the slice without considering the index
  for _,nb := range x {
    //doubling each value
   n:= nb*2
    //storing each value in a new slice
    new_a=append(new_a,n)
  }
  //returning the new slice
  return new_a
}
