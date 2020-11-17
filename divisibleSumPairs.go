package main
  
import ( 
        "fmt"
)

func main() {
    arr := [] int {1, 3, 2, 6, 1, 2}
    fmt.Println(arr[1:])
    fmt.Println(divisibleSumPairs(6, 3, arr))
}

func divisibleSumPairs(n int, k int, arr []int) int{
    var result [n]int
    for i := 0; i < n; i++ {
        result[i] = arr
    }
}

func sum(arr []int) int {
    result := 0
    for _, v := range arr {
        result += v
    }
    return result
}
