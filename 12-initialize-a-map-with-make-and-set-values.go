//Create a map with make function. Initialize the map to store string keys (names of people)
// and integer values (their ages), then add 3 entries to it.
package main
import "fmt"

func main(){
  people := make(map[string] int)
  people["Alice"] = 28
  people["Bob"] = 35
  people["Carol"] = 27

  fmt.Println(people)
}
