//write a function that takes a map of map[string]int and 
//returns the sum of all values in the map.
package main
import "fmt"

func sumValues(m map[string]int)int {
  sum := 0
  for _, value := range m {
    sum += value
  }
  return sum
}

func main(){
  m := map[string]int{"a":10, "b":20, "c":30}
  total := sumValues(m)
  fmt.Println("sum:", total)
}
