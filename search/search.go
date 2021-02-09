package search

/**
 * 線形探索
 */
func linearSearch(str string, strs []string, length int) int{
    for i := 0; i < length ; i++ {
        if str == strs[i] {
            return i
        }
    }
    return -1
}