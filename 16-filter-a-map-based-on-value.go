// write a function that filters a map of string to int, 
// returning only the entries where the value is greater than a given threshold.
package main
import "fmt"

func filterMapByValue(m map[string]int, threshold int) map[string]int {
  filtered := make(map[string]int)
  for k, v := range m {
    if v  > threshold {
      filtered[k] = v
    }
  }
  return filtered
}

func main(){
  fruits := map[string]int{"apple":3, "banana":2, "orange":8}
  fmt.Println("Initial fruits:", fruits)
  threshold := 3
  fmt.Println("threshold >", threshold)
  filtered := filterMapByValue(fruits, threshold)
  fmt.Println("Current fruits:", filtered)
}
