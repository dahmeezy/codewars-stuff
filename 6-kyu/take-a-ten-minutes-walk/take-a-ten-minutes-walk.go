package kata
​
​
func IsValidWalk(walk []rune) bool {
  
  mp:=make(map[rune]int)
  
//   populate map with directions
  for _,direction:=range walk{
    mp[direction]++
  }
  for _ = range mp{
//     checks that exactly 10 directions are passed and count of oppositr directions are the same
    if len(walk)==10 && mp['n']==mp['s'] && mp['e']==mp['w']{
      return true
    }
  }
  return false
}
​