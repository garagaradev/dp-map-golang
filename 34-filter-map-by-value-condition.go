//write a function that filters out entries in a map of map[string]int where 
//the value is greater than a given threshold.
package main
import "fmt"

func filterByValue(m map[string]int, threshold int) map[string]int {
  result := make(map[string]int)
  for key, value := range m {
      if value > threshold {
      result[key] = value
    }
  }
  return result
}

func main(){
  m := map[string]int{"a":10, "b":20, "c":5, "d":15}
  filtered:=filterByValue(m, 10)
  fmt.Println(filtered)
}
