package kata
​
import (
"strings")
​
func ToJadenCase(str string) string {
  var res strings.Builder
  if str==""{
    return " "
  } else{
    words:=strings.Fields(str)
    for index,word:=range words{
      for i,c :=range word{
        if i==0 && (c>='a'&&c<='z'){
          res.WriteRune(rune(int(c)-32))
          
        } else if i >0&&(c>='A'&&c<='Z'){
          res.WriteRune(rune(int(c)+32))
        }else{
          res.WriteRune(c)
        }
        
      }
      if index==len(words)-1{
        continue
      }else{
        res.WriteString(" ")
      }
      
    }
  }  
  return res.String() // your code here...
}