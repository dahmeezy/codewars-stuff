package kata
​
import(
"unicode"
"strings")
​
​
func bubbleSortC(chars []rune) []rune {
  for i := 0; i < len(chars); i++ {
​
    for i, c := range chars {
      temp := chars[i]
      if i < len(chars)-1 && c > chars[i+1] {
        chars[i] = chars[i+1]
        chars[i+1] = temp
      }
    }
  }
  return chars
}
​
func FindChildren(dancingBrigade string) string{
  
  var res strings.Builder
​
//   get the capitals
  var up strings.Builder
  for i, c := range dancingBrigade {
    if unicode.IsUpper(c) {
      up.WriteByte(dancingBrigade[i])
    }
  }
//   cast capitals to rune and uppercase them
  runes := []rune(up.String())
  runed:=bubbleSortC(runes)
​
//   find their children, store them
  for j, u := range runed {
    res.WriteRune(runed[j])
    for i, c := range dancingBrigade {
      if c==unicode.ToLower(u){
        res.WriteByte(dancingBrigade[i])
      }
​
    }
  }
  return res.String()
}
​