//delete an element from the map
package main
import "fmt"

func main(){
  fruits := map[string]int{
    "apple":5,
    "banana":2,
    "orange":8,
  }
  fmt.Println("Initial fruits:", fruits)
  delete(fruits, "banana")
  fmt.Println("Current fruits:", fruits)
}
