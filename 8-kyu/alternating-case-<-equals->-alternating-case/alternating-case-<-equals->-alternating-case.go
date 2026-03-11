package kata
​
func ToAlternatingCase(str string) string {
  
  res:= ""
  
  for _,c := range str{
    
    if c >= 'a' && c <= 'z' {
      
      v:= c - 32
      
      res += string(v)
      
    } else if c >= 'A' && c <= 'Z' {
      
      v:= c + 32
      
      res += string(v)
      
    } else {
      
      res += string(c)
      
    }
    
  }
  return res
  }
  
​