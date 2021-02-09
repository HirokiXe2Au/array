package search

func LinearSearchInt(value int, values []int, length int) int{
  for i := 0; i < length ; i++ {
      if  value == values[i] {
          return i
      }
  }
  return -1
}

func LinearSearchString(str string, strs []string, length int) int{
  for i := 0; i < length ; i++ {
      if  str == strs[i] {
          return i
      }
  }
  return -1
}