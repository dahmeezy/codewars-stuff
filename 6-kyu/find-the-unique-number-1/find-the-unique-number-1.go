package kata
​
func FindUniq(arr []float32) float32 {
  // Do the magic
  mp:=make(map [float32]int)
//   store in a map
  for _,v:=range arr{
    mp[v]++
  }
  var diff float32
//   check for value that appears exactly once and return
  for i,count:=range mp {
    if count ==1{
      diff=i
    }
  }
  return diff
}