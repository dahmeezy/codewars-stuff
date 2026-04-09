package kata
​
func Factorial(n int) int {
  // put your code here
  no:=1
  for i:=n;i>0;i--{
    no*=i
  }
  return no
}