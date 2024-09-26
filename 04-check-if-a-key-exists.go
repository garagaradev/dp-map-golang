//check if a key exists in the map
package main
import "fmt"

func main(){
  fruits := map[string]int{
    "apple":2,
    "banana":8,
    "orange":7,
  }
  fmt.Println("the fruits:", fruits)

  if value, exists := fruits["grape"]; exists {
    fmt.Println("Grape exists with value:", value)
  } else {
    fmt.Println("Grape doesn't exist")
  }
}
