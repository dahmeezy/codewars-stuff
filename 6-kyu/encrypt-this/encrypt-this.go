package kata
​
import (
"strings"
"strconv")
​
func EncryptThis(text string) string {
  
  // Implement me! :)  
  if len(text)<1{
    return text
  }
  var res strings.Builder
  var finalRes strings.Builder
  
  words:=strings.Fields(text)
  
  
  for j,word:=range words{
  
  runed:=[]rune(word)
    if len(runed)==1{
      res.WriteString(strconv.Itoa(int(runed[0])))
    }else if len(runed)==2{
      res.WriteString(strconv.Itoa(int(runed[0])))
      res.WriteRune(runed[1])
    }else{
      
  asciiOfFirst:=int(runed[0])
  
  res.WriteString(strconv.Itoa(asciiOfFirst))
  
  res.WriteRune(runed[len(word)-1])
  
  for i:=2;i<=len(runed)-2;i++{
    res.WriteRune(runed[i])
  }
  res.WriteRune(runed[1])
    
    }
    if j !=len(words)-1{
      finalRes.WriteString(res.String())
    finalRes.WriteString(" ")
    }else{
      finalRes.WriteString(res.String())
      }
    res.Reset()
  
  }
  return finalRes.String()
}