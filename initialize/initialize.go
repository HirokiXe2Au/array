package initialize

func LinearNumber(start int, end int) []int{
  abs := func(x int) int{
    if x < 0 {
      x = -x
    }
    return x
  }

  var length int = abs(start - end)
  list := make([]int, length)
  if end > start {
    for i := 0; i < length ; i++ {
      list[i] = start + i 
    }
  }else{
    for i := 0; i < length ; i++ {
      list[i] = start - i
    }
  }
  return list
}

func LinearArray(start int, end int, ary []string) []string{
  var convert func(n int, ary []string) string
  convert = func(n int, ary []string) string{
    var str string
    i := n % len(ary)
    r := n / len(ary)
    if r > 0{
      str = convert(r-1, ary)
    }
    return  str + ary[i]
  }

  abs := func(x int) int{
    if x < 0 {
      x = -x
    }
    return x
  }

  var length int = abs(start - end)
  strs := make([]string, length)
  if end > start {
    for i := 0; i < length ; i++ {
      strs[i] = convert(abs(start + i), ary)
    }
  }else{
    for i := 0; i < length ; i++ {
      strs[i] = convert(abs(start - i), ary)
    }
  }
  return strs
}

func LinearUpperCapital(start int, end int) []string{
  tmp := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
  return LinearArray(start, end, tmp)
}

func LinearLowerCapital(start int, end int) []string{
  tmp := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
  return LinearArray(start, end, tmp)
}