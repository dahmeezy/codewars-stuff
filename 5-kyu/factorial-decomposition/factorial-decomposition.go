  if nb<2{
    return 2
  }
  for{
    nb++
    if isPrime(nb){
      return nb
    }
  }
}
​
func Decomp(n int) string {
    // your code
  mp:=make(map[int]int)
  sl:=make([]int, n-1)
  num:=2
  
  var res strings.Builder
  
  for i:=0;i<len(sl);i++{
    sl[i]=num
    num++
  }
  
  for i:=range sl{
    temp:=sl[i]
    prime:=2
    for temp > 1{
    if temp%prime==0{
//       populate map if prime is a factor of current value
      mp[prime]++
    }else{
//       prime becomes the next one
      prime=findNextPrime(prime)
      continue
    }
    temp=temp/prime
    }
  }
  var dum []int
  for _,v:=range sl{
    if isPrime(v){
      dum=append(dum,v)
      }
  }
    for i,v:=range dum{
      res.WriteString(strconv.Itoa(v))
      if mp[v]!=1{
      res.WriteString("^")
      res.WriteString(strconv.Itoa(mp[v]))
      }
      if i!=len(dum)-1{
      res.WriteString(" * ")
        continue
      }
    }
  return res.String()
}