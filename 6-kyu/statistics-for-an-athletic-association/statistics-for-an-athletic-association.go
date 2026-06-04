​
  // use bubble sort to sort slice
  for i := valLength - 1; i > 0; i-- {
    for j := range val {
      temp := val[j]
      if j < valLength-1 && val[j] > val[j+1] {
        val[j] = val[j+1]
        val[j+1] = temp
      }
    }
  }
​
  // get the range in int
  r := val[valLength-1] - val[0]
​
  // sum the values
  var sumOfVals int
  for _, v := range val {
    sumOfVals += v
  }
  // use sum to get average value
  a := sumOfVals / valLength
​
  // median
  md := 0
  if valLength%2 == 0 {
    v1 := val[valLength/2-1]
    v2 := val[valLength/2]
    md = (v1 + v2) / 2
  } else {
    md = val[valLength/2]
  }
​
  rangeInHrs = toHours(r)
​
  avg = toHours(a)
​
  median = toHours(md)
​
  res := "Range: " + rangeInHrs + " Average: " + avg + " Median: " + median
​
  return res
​
}
​
// an helper function to put the range, mean and median in required format
func toHours(val int) string {
  h := val / 3600
  m := (val % 3600) / 60
  s := (val % 3600) % 60
​
  hAsString := strconv.Itoa(h)
  mAsString := strconv.Itoa(m)
  sAsString := strconv.Itoa(s)
​
  if h >= 0 && h <= 9 {
    hAsString = "0" + hAsString
  }
  if m >= 0 && m <= 9 {
    mAsString = "0" + mAsString
  }
  if s >= 0 && s <= 9 {
    sAsString = "0" + sAsString
  }
​
  res := hAsString + "|" + mAsString + "|" + sAsString
​
  return res
​
}