package kata
​
func GetGrade(a,b,c int) rune {
    // Code here
  avg:=(a+b+c)/3
 if avg>=90 && avg <=100{
   return 'A'
 } else if avg >=80 && avg<90{
   return 'B'
 } else if avg >=70 && avg<80{
   return 'C'
 }else if avg >=60 && avg<70{
   return 'D'
 } 
  return 'F'
}