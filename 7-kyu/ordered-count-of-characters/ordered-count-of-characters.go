package orderedcount
​
import "strings"
​
// Use the preloaded Tuple struct as return type
//type Tuple struct {
 //   Char  rune
//    Count int }
​
func OrderedCount(text string) []Tuple {
  // Implement me! :)
  var res []Tuple
striped:=""
  
  mp:=make(map[rune]int)
  for _,c:=range text{
    mp[c]++
}
  
  for _,c:=range text{
    if strings.ContainsAny(striped,string(c)){
      continue
    }
  striped+=string(c)}
for _,c:=range striped{
  t:=Tuple{Char:c,Count:mp[c]}
  res=append(res,t)
}
  if res==nil{
return []Tuple{}
  }
  return res
}