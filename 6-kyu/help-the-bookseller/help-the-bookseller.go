package kata
​
import ("strings"
        "strconv"
        )
​
func StockList(listArt []string, listCat []string) string {
  if len(listArt) == 0 || len(listCat) == 0 {
    return ""
}
  
  // your code
  mp := make(map[string]int)
​
  var code string
  var unit int
  
​
  var res string
​
  for _, data := range listArt {
    part := strings.Split(data, " ")
    code = part[0]
    unit, _ = strconv.Atoi(part[1])
    for _, l := range listCat {
      mp[l] += 0
      if l == string(code[0]) {
        mp[l] += unit
      }
​
    }
​
  }
​
  for i, cat := range listCat {
    res += "(" + cat + " : " + strconv.Itoa(mp[cat]) + ")"
    if i < len(listCat)-1 {
      res += " - "
    }
  }
  return res
​
​
}