//write a function that stores data of different types (int, string, bool)
//in a map using map[string]interface{}.Then print the values.
package main
import "fmt"

func main(){
  data := map[string]interface{}{
    "name":"Jack",
    "age":35,
    "active":true,
  }

  for key, value := range data {
    fmt.Printf("%s -> %v\n", key, value)
  }
}
