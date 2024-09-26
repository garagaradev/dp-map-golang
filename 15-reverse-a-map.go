// Write a function that reverses a map. The keys should become the values and the values should become the keys.
// Assume the values are unique and can be used as keys.
package main
import "fmt"

func reverseMap(m map[string]int) map[int]string {
  reversed := make(map[int]string)
  for k, v := range m {
    reversed[v] = k
  }
  return reversed
}

func main(){
  fruits := map[string]int{"apple":5,"banana":2,"orange":8}
  fmt.Println("Initial form:",fruits)
  reversed := reverseMap(fruits)
  fmt.Println("Current form:", reversed)
}
