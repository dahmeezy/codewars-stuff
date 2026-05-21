package kata
​
func Solve(s string) []int {
  //..
  var uc int
  var lc int
  var n int
  var sc int
  
  var res []int
  
  for _,c:=range s{
    if c>='A'&& c<='Z'{
      uc++
    } else if c>='a'&&c<='z'{
      lc++
    }else if c>='0'&&c<='9'{
      n++
    } else{
      sc++
    }
  }
  res=append(res,uc)
  res=append(res,lc)
  res=append(res,n)
  res=append(res,sc)
  
  return res
}