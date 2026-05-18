package kata
​
func HighestRank(nums []int) int {
  mp:=make(map[int]int)
  
  for _,n := range nums{
    mp[n]++
  }
  h:=0
  highestK := 0
  
  for k,v:=range mp{
    if v > h{
      h=v
      highestK = k
     }
    if v == h &&  k > highestK {
      highestK = k
    }
    }
      return highestK
  }
​
​