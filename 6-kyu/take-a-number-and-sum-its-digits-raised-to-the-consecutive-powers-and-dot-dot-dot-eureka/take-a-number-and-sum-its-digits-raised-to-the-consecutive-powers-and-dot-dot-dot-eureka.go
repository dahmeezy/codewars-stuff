package kata
​
import (
"strconv"
"math")
​
​
func SumDigPow(a, b uint64) (res []uint64) {
  for i:=a; i<=b; i++{
    if i>=1 && i <=9{
     res=append(res, i)
    }else{
      n:=strconv.FormatUint(i,10)
        hm:=0
      for j,c:=range n{
        nu,_:=strconv.Atoi(string(c))
        hm+=int(math.Pow(float64(nu),float64(j+1)))
        if hm==int(i){
        res=append(res, uint64(hm))
     }
        }
    }
}
  return res
}
​