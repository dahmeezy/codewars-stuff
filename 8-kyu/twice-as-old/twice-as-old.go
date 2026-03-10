package kata
​
//import "math"
​
func TwiceAsOld(dyear, syear int) int { 
  // Implement me
  if istwice(dyear, syear)==true{
    gap:=dyear-(syear*2)
    return gap;
  } else if istwice(dyear,syear)==false{
    gap:=(syear*2)-dyear
    return gap
  }
  return 0
}
​
func istwice(d, s int) bool{
  if d<2*s{
    return false
  }
  return true
}