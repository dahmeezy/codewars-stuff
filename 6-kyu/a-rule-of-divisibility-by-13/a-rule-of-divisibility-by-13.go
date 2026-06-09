package kata
​
import "strconv"
​
func Thirt(n int) int {
  // your code
  multipliers := []int{1, 10, 9, 12, 3, 4}
​
  prevVal := -1
​
  stationaryNum := n
​
//   the loop should continue as long as the new value is not same as the previous one
  for prevVal != stationaryNum {
​
    prevVal = stationaryNum
​
//     convert current stationary value to string in order to loop through them
    stringedNum := strconv.Itoa(stationaryNum)
​
    count := 0
​
    summedDigit := 0
​
    for i := len(stringedNum) - 1; i >= 0; i-- {
      digit, _ := strconv.Atoi(string(stringedNum[i]))
//       this also helps when length of digit exceeds the length of multipliers
      multiplier := multipliers[count%6]
      summedDigit += digit * multiplier
      count++
    }
    stationaryNum = summedDigit
​
  }
  return stationaryNum
​
}