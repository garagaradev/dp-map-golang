// create a simple map
package main
import "fmt"

func main(){
  fruits := map[string]int{
    "apple": 5,
    "banana": 2,
    "orange": 8,
  }
  fmt.Println("The amount of banana:",fruits["banana"])
}
