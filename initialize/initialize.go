package initialize

func LinearNumber(start int, end int) []int{
  var diff int = start - end
  var length int
  if diff >= 0 {
    length = diff
  }else{
    length = -diff
  }

  list := make([]int, length)
  for i := 0; i < length ; i++ {
    if end >= start {
      // 正の場合
      list[i] = start + i 
    }else{
      // 負の場合
      list[i] = start - i
    }
  }
  return list
}

func convert(n int, ary []string) string{
  var str string
  i := n % len(ary)
  r := n / len(ary)
  if r > 0{
    str = convert(r-1, ary)
  }
  return  str + ary[i]
}

func LinearArray(start int, end int, ary []string,) []string{
  var diff int = start - end
  var length int
  if diff >= 0 {
    length = diff
  }else{
    length = -diff
  }

  list := make([]string, length)
  var num int
  if end >= start {
    for i := 0; i < length ; i++ {
      num = start + i
      if num < 0 {
        num = -num
      }
      list[i] = convert(num, ary)
    }
  }else{
    for i := 0; i < length ; i++ {
      num = start - i
      if num < 0 {
        num = -num
      }
      list[i] = convert(num, ary)
    }
  }
  return list
}

func LinearUpperCapital(start int, end int) []string{
  tmp := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
  return LinearArray(start, end, tmp)
}

func LinearLowerCapital(start int, end int) []string{
  tmp := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
  return LinearArray(start, end, tmp)
}