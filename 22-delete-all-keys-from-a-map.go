// writes a function that clears all entries in a map.
// the map has string keys and integer values.
package main
import "fmt"

func clearMap(m map[string]int) {
  for key := range m {
    delete(m, key)
  }
}

func main(){
  numbers := map[string]int{"one":1, "two":2, "three":3}
  fmt.Println("Initial numbers:", numbers)
  clearMap(numbers)
  fmt.Println("Current number:", numbers)

}
