# Array Package
## initialize
### Create linear numbers
Create array to numbers in first argument to second argument order.
```
// Get array to []int{10 9 8 7 6 5 4 3 2 1 0 -1 -2 -3 -4 -5 -6 -7 -8 -9}
numbers := LinearNumber(10, -10)
```
### Create linear strings
Create array to strings in first argument to second argument order.
```
// Get array to []string{"K", "J", "I", "H", "G", "F", "E", "D", "C", "B", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
strs := LinearUpperCapital(10, -10)

// Get array to []string{"x", "y", "z", "aa", "ab", "ac", "ad"}
strs := LinearLowerCapital(23, 30)
```
## search
### linear search
Use Linear algorithm to search.
```
// Search 3 in arr, result's value is 7
arr := []int{10 9 8 7 6 5 4 3 2 1 0 -1 -2 -3 -4 -5 -6 -7 -8 -9}
result := LinearSearchInt(3, arr, 20)

// Search "J" in arr, result's value is 1
arr := []string{"K", "J", "I", "H", "G", "F", "E", "D", "C", "B", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
result := LinearSearchString("J", arr, 20)
```