package kata
​
func IsTriangle(a, b, c int) bool {
  
  if a<=0||b<=0||c<=0{
    return false
  }
  
  if a<b+c&&b<a+c&&c<a+b{
    return true
  }
  return false
​
}