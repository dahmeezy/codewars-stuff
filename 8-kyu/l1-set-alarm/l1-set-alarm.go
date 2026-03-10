package kata
​
​
func SetAlarm(emp, vac bool) bool {
  if emp && vac{
    return false
  } else if emp && !vac{
    return true
  } else if !emp && vac{
    return false
  } 
  return false
}