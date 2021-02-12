package initialize

const m = 32 << (^uint(0) >> 63) - 1
func abs(x int) int{
  if x < 0 {
    x = -x
  }
  return (x ^ (x >> m)) - (x >> m)
}

func LinearNumber(start int, end int) []int{
  var length int = abs(start - end)
  ary := make([]int, length)
  if end > start {
    for i := 0; i < length ; i++ {
      ary[i] = start + i 
    }
  }else{
    for i := 0; i < length ; i++ {
      ary[i] = start - i
    }
  }
  return ary
}

func LinearArray(start int, end int, strs []string) []string{
  var convert func(n int) string
  var l int = len(strs)
  convert = func(n int) string{
    var str string
    i := n % l
    r := n / l
    if r > 0{
      str = convert(r-1)
    }
    return  str + strs[i]
  }

  var length int = abs(start - end)
  ary := make([]string, length)
  if end > start {
    for i := 0; i < length ; i++ {
      ary[i] = convert(abs(start + i))
    }
  }else{
    for i := 0; i < length ; i++ {
      ary[i] = convert(abs(start - i))
    }
  }
  return ary
}

func LinearUpperCapital(start int, end int) []string{
  tmp := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
  return LinearArray(start, end, tmp)
}

func LinearLowerCapital(start int, end int) []string{
  tmp := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
  return LinearArray(start, end, tmp)
}