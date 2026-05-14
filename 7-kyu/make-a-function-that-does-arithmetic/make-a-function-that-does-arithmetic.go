package kata
​
func Arithmetic(a int, b int, operator string) int{
  //your code here
    if operator =="add"{
      return (a + b)
    } else if operator =="multiply"{
      return (a * b)
    }else if operator =="subtract"{
      return (a - b)
    }else if operator =="divide"{
      return (a / b)
    } 
  return 0
    
}