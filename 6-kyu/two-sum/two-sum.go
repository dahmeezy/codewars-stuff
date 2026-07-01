package kata
​
func TwoSum(numbers []int, target int) (res [2]int) {
//   map whose key is the number and its value is the index of that number
  mp:=make(map[int]int)
  for i,v:=range numbers{
//     store the complement
    comp:=target-v
//     checkif complement has been seen before(exists inside map)
    val,ok:=mp[comp]
    if ok{
      res[0]=val
      res[1]=i
    }else{
//       store the current number and its index in map
      mp[v]=i
    }
  }
  return res
}