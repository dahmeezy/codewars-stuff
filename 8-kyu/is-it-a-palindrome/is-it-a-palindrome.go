package kata
​
func IsPalindrome(str string) bool {
  // change the string to lower
  runed:=[]rune(str)
    for i,char:=range str{
      if char>='A'&&char<='Z'{
        runed[i]=char+32
      }
    }
  runes:=string(runed)
  var pali string
  for i:=len(runes)-1;i>=0;i--{
    pali+=string(runes[i])
  }
  if pali== runes{
    return true
  }
  return false
}