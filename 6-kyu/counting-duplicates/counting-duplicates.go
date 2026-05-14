package kata
​
import "strings"
​
func duplicate_count(s1 string) int {
  //your code goes here
//   normalize the string
  s:=strings.ToLower(s1)
//   make a map for the characters and their frequencies
  mp:=make(map[rune]int)
//   create a variable to store an int
  count:=0
//   loop through s and populate the map with its value
  for _,c:=range s{
    mp[c]++
  }
// check for one that has duplicate and record
  for _,v:=range mp{
    if v>1{
      count++
    }
  }
  
  return count//Instead of returning '1', you have to return integer 'i' as answer of solution.
}
​
​
​
​