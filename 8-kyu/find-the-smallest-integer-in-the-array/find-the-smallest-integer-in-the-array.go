package kata
​
func SmallestIntegerFinder(numbers []int) int {
  //catching cases where the slice might be empty
  if len(numbers)==0{
    return 0
  }
  //using the first integer in the slice as a point of reference
  first:=numbers[0]
  //a for loop to iterate all over
  for i:=1;i<=len(numbers)-1;i++{
    //set current number to number of index i
    current:=numbers[i]
    // if the first number is greater, the current number becomes the new point of reference
    if first > current{
      first=current
      continue
      //if not the first one keep being the point of reference
    } else {
      continue
    }
  }
  // prints out the least value
  return first // your code here
}