package kata
​
func HighestRank(nums []int) int {
  mp:=make(map[int]int)
  
//   populate the map with the values and their frequencies
  for _,n:=range nums{
    mp[n]++
  }
//   variables for the highest frequency amd the mode(the number itself)
  hfreq:=0
  mode:=0
  for k,v:=range mp{
//     get the highest frequency and mode
    if v>hfreq{
      hfreq=v
      mode=k
    }
//   If there is a tie for most frequent number, return the largest number among them.
    if v==hfreq && k>mode{
      mode =k
    }
  }
    return mode
}
​