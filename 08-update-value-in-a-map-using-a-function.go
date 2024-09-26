// write a function that takes a map of string to int and adds 10 to each value
package main
import "fmt"

func addTenToEachValue(m map[string]int) {
  for key := range m {
    m[key] += 10
  }
}

func main(){
  fruits := map[string] int {
    "apple":5,
    "banana":2,
    "orange":8,
  }
  fmt.Println("Initial fruits:", fruits)
  addTenToEachValue(fruits)
  fmt.Println("Current fruits:", fruits)
}
