package kata
​
import ("strings"
       "unicode"
       "strconv")
​
func Order(sentence string) string {
  
  words:=strings.Split(sentence," ")
  
  res:=make([]string,len(words))
  
  for _,word:=range words{
    for _,c:=range word{
      if unicode.IsDigit(c){
        
        d, _:=strconv.Atoi(string(c))
        res[d-1]=word
      }
    }
  }
  return strings.Join(res," ")
}
​