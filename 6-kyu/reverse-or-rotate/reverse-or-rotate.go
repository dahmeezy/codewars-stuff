  for i := 0; i < lenstr; i += ch {
​
    end := i + ch
​
    if end > len(s) {
      end = len(s)
    }
    chunk := s[i:end]
    if len(chunk) == ch {
      box = append(box, chunk)
    }
​
  }
  
// check sum of digits in each chunk
  for i, num := range box {
    var r string
    sum := 0
    for _, d := range num {
​
      val, _ := strconv.Atoi(string(d))
      sum += val
​
    }
//     if its even, reverse
    if sum%2 == 0 {
      for i := ch - 1; i >= 0; i-- {
        r += string(num[i])
      }
      box[i] = r
      
//       if not rotate
    } else {
​
      for ind, d := range num {
        if ind != 0 {
          r += string(d)
        }
      }
      r += string(num[0])
      box[i] = r
    }
  }
//   join what u have and return
  return strings.Join(box,"")
​
}
​