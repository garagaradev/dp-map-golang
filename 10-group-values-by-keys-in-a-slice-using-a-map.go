// given a slice of integers, group them by their remainder when divided by 3 using a map
package main
import "fmt"

func groupByRemainder(nums []int) map[int][]int{
    result := make(map[int][]int)
    for _, num := range nums {
    remainder := num % 3
    result[remainder] = append(result[remainder], num)
  }
  return result
}

func main(){
  nums := []int{1,2,3,4,5,6,7,8,9}
  grouped := groupByRemainder(nums)
  fmt.Println(grouped)
}
