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

func LinearArray(start int, end int, ary []string, len int) []string{
  var diff int = start - end
  var length int
  if diff >= 0 {
    length = diff
  }else{
    length = -diff
  }

  list := make([]string, length)
  var num int
  for i := 0; i < length ; i++ {
    if end >= start {
      num = (start + i) % len
    }else{
      num = (start - i) % len
    }

    if num < 0 {
      num = -num
    }
    list[i] = ary[num]
  }
  return list
}

func LinearUpperCapital(start int, end int) []string{
  tmp := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
  return LinearArray(start, end, tmp, 26)
}

func LinearLowerCapital(start int, end int) []string{
  tmp := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
  return LinearArray(start, end, tmp, 26)
}