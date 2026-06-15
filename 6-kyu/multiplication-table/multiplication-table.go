package multiplicationtable
​
func MultiplicationTable(size int) [][]int {
  // Implement me! :)
  var v int
  var multiTable [][]int
  for c := 0;c<size;c++ {
//     creates a new slice each time
    var res []int
    for r := 0;r<size;r++{
//       row times column
      v = (c + 1) * (r + 1)
//      
      res = append(res, v)
    }
    multiTable = append(multiTable, res)
  }
  return multiTable
}