package search

/**
 * 線形探索
 */
func linearSearch(str interface{}, strs []interface{}, length int) int{
  for i := 0; i < length ; i++ {
      if  str.(type) == strs[i].(type) && str == strs[i] {
          return i
      }
  }
  return -1
}