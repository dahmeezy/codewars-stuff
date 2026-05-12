package kata
​
import (
"strings"
"unicode")
​
func ToNato(words string) string {
// the NATO map[string]string is preloaded for you
// NATO["A"] == "Alfa", etc
  
//   strip string of all spaces
 w:=strings.ReplaceAll(strings.ToUpper(words)," ","")
//   create a container
  var res strings.Builder
//   loop through and store the corresponding value in the map
  for i,c := range w{
    res.WriteString(NATO[string(c)])
//     store like that if its a punctuation
    if unicode.IsPunct(c){
      res.WriteRune(c)
    }
//     dont add space after the last word or chat in the string
    if i != len(w)-1{
      res.WriteString(" ")
    }
  
  }
    return res.String()
}